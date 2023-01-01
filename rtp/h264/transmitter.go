package h264

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "/usr/local/usr/include/video_reader.h"
import "C"
import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"time"
	"tmss/media"
	"tmss/rtp/parser"
	"tmss/rtsp/headers"
	"unsafe"
)

const DefaultFragmentationType = 28
const (
	GetNextPacket = iota
	StopStream    = iota
	Play          = iota
	Pause         = iota
	Idle          = iota
	Seek
)

type MediaStreamer interface {
	Play(timeRange headers.Range)
	Pause(timeRange headers.Range)
	HandleRtcp()
	HandleRtp()
}
type Streamer struct {
	mediaId        string
	mediaRecord    media.Media
	buffer         *Buffer
	bufferCommand  chan BufferCommand
	rtpConn        net.PacketConn
	rtcpConn       net.PacketConn
	OutByteRate    int
	streamID       int
	PayloadType    byte
	SequenceNumber int32
	SSRC           int32
	tsIncrement    uint32
	rtpTimestamp   uint32
	ClientAddr     net.Addr
	CodecHeader    Header
}

func Init(mediaRecord media.Media, streamId int, rtpConn net.PacketConn, rtcpConn net.PacketConn) (*Streamer, error) {
	streamer := &Streamer{
		rtpConn:     rtpConn,
		rtcpConn:    rtcpConn,
		streamID:    streamId,
		mediaRecord: mediaRecord,
		PayloadType: byte(mediaRecord.Streams[streamId].PayloadType),
	}
	decodedHeaderBytes, err := base64.StdEncoding.DecodeString(mediaRecord.Streams[streamId].HeaderB64)
	if err != nil {
		return nil, err
	}
	streamer.CodecHeader = ParseAvccHeader(decodedHeaderBytes)
	streamer.OutByteRate = int(math.Ceil(float64(mediaRecord.Streams[streamId].BitRate) / float64(BitsInByte)))
	fmt.Printf("byte rate: %v\n", streamer.OutByteRate)
	buffSizeByte := C.int(BuffSizeFactor * streamer.OutByteRate)
	// TODO change Stream.Path to Stream.Name
	streamPath := C.CString(media.GetStreamPath(mediaRecord.Streams[streamId].Path))
	fmt.Printf("stream path: %v\n", streamPath)
	mediaBuffer := C.init_media_buffer(streamPath, buffSizeByte)
	if mediaBuffer == nil {
		return nil, errors.New("failed to create media buffer")
	}
	streamer.buffer = &Buffer{
		CurrentBuffer: mediaBuffer,
		buffChan:      make(chan int),
		isBuffering:   false,
	}
	return streamer, nil
}

type BufferCommand struct {
	position headers.Range
	command  int
}

func (s *Streamer) Play(timeRange headers.Range) {
	//TODO implement me
	panic("implement me")
}

func (s *Streamer) Pause(timeRange headers.Range) {
	panic("implement me")
}

func (s *Streamer) HandleRtp() {
	go func() {
		fmt.Printf("trying to send rtp packets\n")
		buff := make([]byte, MTU)
		nBytesRead, a, err := s.rtpConn.ReadFrom(buff)
		if err != nil {
			return
		}

		s.ClientAddr = a
		stream := s.mediaRecord.Streams[s.streamID]
		_ = nBytesRead
		_ = a
		_ = stream
		lastCommand := BufferCommand{
			command: Play,
		}
		t0 := 0
		maxByteOut := s.OutByteRate
		remainingSize := maxByteOut
		go s.buffer.BufferUp()

		// send SPS and PPS

		s.SendPacket(s.CodecHeader.SPS)
		fmt.Printf("%v\n", s.CodecHeader.SPS)
		s.SendPacket(s.CodecHeader.PPS)

		for {

			select {
			case newCommand := <-s.bufferCommand:
				lastCommand = newCommand
			default:
				switch lastCommand.command {
				case Play:
					{
						println("sending AVPackets.....")
						if t0 <= 0 || time.Now().Sub(time.Unix(0, int64(t0))) > time.Second {
							t0 = time.Now().Nanosecond()
							remainingSize = maxByteOut
						}
						/*avPacket := s.buffer.ReadNextPacket(true)

						fmt.Printf("sending packet; pts: %v, size: %v\n", avPacket.pts, avPacket.size)

						if int(avPacket.size) > remainingSize {
							continue
						}*/
						avPacket := s.buffer.ReadNextPacket(false)
						if avPacket == nil {
							return
						}
						encodedPacket := C.GoBytes(unsafe.Pointer(avPacket.data), C.int(avPacket.size))

						fmt.Printf("packets: %v\n", encodedPacket[:20])
						time.Sleep(time.Millisecond * 20)
						nalUnits := AVCCToNalU(s.CodecHeader, encodedPacket)

						for _, nalUnit := range nalUnits {
							s.SendPacket(nalUnit)
						}
						//TODO send packet
						remainingSize -= int(avPacket.size)
					}
				case Pause:
					{
						lastCommand = <-s.bufferCommand
					}
				case Seek:
					{

					}
				}
			}
		}
	}()
}

func (s *Streamer) HandleRtcp() {
	go func() {
		for {
			buff := make([]byte, MTU)
			n, from, err := s.rtcpConn.ReadFrom(buff)
			fmt.Printf("received msg from: %s\n%s\n", from, string(buff[:n]))
			if err != nil {
				return
			}

		}
	}()

}

func (s *Streamer) SendPacket(data []byte) {
	var fragments [][]byte
	if len(data) > MTU {
		fragments = packetizeNalUnit(data, DefaultFragmentationType, MTU)
	} else {
		fragments = [][]byte{data}
	}
	s.rtpTimestamp += s.tsIncrement
	for _, f := range fragments {
		rtpHeader := parser.Header{
			Version:        parser.RtpVersion,
			PayloadType:    s.PayloadType,
			SequenceNumber: uint16(s.GetNextSeqNumber()),
			Timestamp:      s.rtpTimestamp,
			SSRC:           uint32(s.SSRC),
		}
		rtpPacket := parser.SerializeRTPPacket(rtpHeader, f)
		n, err := s.rtpConn.WriteTo(rtpPacket, s.ClientAddr)
		if err != nil {
			log.Fatalf("failed to send data, err : %v\n", err)
			return
		}
		fmt.Printf("wrote %v to %v\n", n, s.ClientAddr)
	}
	return
}

func (s *Streamer) GetNextSeqNumber() int32 {
	if s.SequenceNumber == 0 {
		s.SequenceNumber = rand.Int31()
	}
	s.SequenceNumber += 1
	return s.SequenceNumber
}
