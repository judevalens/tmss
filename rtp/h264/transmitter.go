package h264

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "/usr/local/usr/include/video_reader.h"
import "C"
import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/pion/rtcp"
	"go.uber.org/zap"
	"log"
	"math"
	"math/rand"
	"net"
	"reflect"
	"sync"
	"time"
	"tmss/media"
	"tmss/misc"
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
	mediaId          string
	mediaRecord      media.Media
	buffer           *Buffer
	bufferCommand    chan BufferCommand
	rtpConn          net.PacketConn
	rtcpConn         net.PacketConn
	OutByteRate      int
	streamID         int
	PayloadType      byte
	SequenceNumber   int32
	SSRC             uint32
	tsIncrement      uint32
	rtpTimestamp     uint32
	ClientAddr       net.Addr
	RtcpClientAddr   net.Addr
	RtcpTimestamp    uint32
	CodecHeader      Header
	logger           zap.Logger
	NumPacketSent    uint32
	lastRtpPacket    time.Time
	lastRtpTimestamp uint32
	SenderMutex      *sync.RWMutex
	packetCount      uint32
}

func Init(mediaRecord media.Media, streamId int, rtpConn net.PacketConn, rtcpConn net.PacketConn) (*Streamer, error) {
	streamer := &Streamer{
		rtpConn:     rtpConn,
		rtcpConn:    rtcpConn,
		streamID:    streamId,
		mediaRecord: mediaRecord,
		PayloadType: byte(mediaRecord.Streams[streamId].PayloadType),
		SenderMutex: &sync.RWMutex{},
	}
	decodedHeaderBytes, err := base64.StdEncoding.DecodeString(mediaRecord.Streams[streamId].HeaderB64)
	if err != nil {
		return nil, err
	}
	streamer.CodecHeader = ParseAvccHeader(decodedHeaderBytes)
	streamer.OutByteRate = int(math.Ceil(float64(mediaRecord.Streams[streamId].BitRate) / float64(BitsInByte)))
	streamer.tsIncrement = uint32(mediaRecord.Streams[streamId].TSIncrement)
	buffSizeByte := C.int(BuffSizeFactor * streamer.OutByteRate)
	// TODO change Stream.Path to Stream.Name
	streamPath := C.CString(media.GetStreamPath(mediaRecord.Streams[streamId].Path))
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
		zap.L().Sugar().Infow("waiting for rtp connection", "stream id", s.streamID)
		buff := make([]byte, MTU)
		_, a, err := s.rtpConn.ReadFrom(buff)
		if err != nil {
			zap.L().Sugar().Errorw("rtp conn is interrupted", "err", err)
			return
		}
		stream := s.mediaRecord.Streams[s.streamID]
		_ = stream
		s.ClientAddr = a
		lastCommand := BufferCommand{
			command: Play,
		}
		t0 := time.Now().UnixNano()
		maxByteOut := s.OutByteRate
		remainingSize := maxByteOut
		go s.buffer.BufferUp()
		i := 0
		// sending SPS and PPS data
		err = s.SendPacket(s.CodecHeader.SPS, 0)
		if err != nil {
			zap.L().Sugar().Fatalw("failed to send SPS data", "stream id", s.streamID, "remote addr", s.ClientAddr, "err", err)
			return
		}
		err = s.SendPacket(s.CodecHeader.PPS, 0)
		if err != nil {
			zap.L().Sugar().Fatalw("failed to send PPS data", "stream id", s.streamID, "remote addr", s.ClientAddr, "err", err)
			return
		}
		for {
			select {
			case newCommand := <-s.bufferCommand:
				lastCommand = newCommand
			default:
				switch lastCommand.command {
				case Play:
					{
						zap.L().Sugar().Debugw("sending packets", "stream id", s.streamID, "client addr", s.ClientAddr)
						if time.Now().Sub(time.Unix(0, int64(t0))) > time.Millisecond*250 {
							remainingSize = maxByteOut
							t0 = time.Now().UnixNano()
						}

						/*avPacket := s.buffer.ReadNextPacket(true)


						if int(avPacket.size) > remainingSize {
							continue
						}*/
						avPacket := s.buffer.ReadNextPacket(false)
						if avPacket == nil {
							zap.L().Sugar().Infow("failed to read packet", "stream id", s.streamID, "client addr", s.ClientAddr)
							return
						}
						encodedPacket := C.GoBytes(unsafe.Pointer(avPacket.data), C.int(avPacket.size))
						//time.Sleep(time.Millisecond * 50)
						nalUnits := AVCCToNalU(s.CodecHeader, encodedPacket)

						for _, nalUnit := range nalUnits {
							err = s.SendPacket(nalUnit, uint32(avPacket.pts)/uint32(1))
							if err != nil {
								zap.L().Sugar().Errorw("failed to send frame", "stream id", s.streamID, "remote addr", s.ClientAddr, "err", err)
								log.Fatal(err)
								return
							}
						}
						//TODO send packet
						remainingSize -= int(avPacket.size)

						/// reminder : tuning this paramter gives better stream and picture quality, optimal i = 60, sleep == 1s
						if i > 60 {
							i = 0
							time.Sleep(time.Millisecond * 1000)
						}
						i++

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
		firstPacket := true
		for {
			buff := make([]byte, 65000)
			n, from, err := s.rtcpConn.ReadFrom(buff)
			fmt.Printf("received msg from: %s\n%s\n", from, string(buff[:n]))
			if err != nil {
				return
			}
			packets, err := rtcp.Unmarshal(buff[:n])
			if err != nil {
				zap.L().Sugar().Errorw("failed to parse rtcp packet...", "err", err)
				continue
			}
			if firstPacket {
				s.RtcpClientAddr = from
				firstPacket = false
				s.SendStats()
			}
			s.ParseRtcpPacket(packets)
		}
	}()
}

func (s *Streamer) SendStats() {
	go func() {
		for {
			s.SenderMutex.RLock()
			now := time.Now()
			elapsed := time.Since(s.lastRtpPacket)
			elapsedToRtp := uint32(elapsed.Seconds()*90000) + s.lastRtpTimestamp
			s.RtcpTimestamp = elapsedToRtp //uint32(now.Unix() * 90000)
			zap.L().Sugar().Infow("rtcp ts", "rtp ts", elapsedToRtp, "now", now.Unix(), "last", s.lastRtpPacket)
			//s.RtcpTimestamp += elapsedToRtp
			senderReport := rtcp.SenderReport{
				PacketCount: s.packetCount,
				SSRC:        s.SSRC,
				NTPTime:     misc.UtcToNtp(now.Unix()),
				RTPTime:     s.RtcpTimestamp,
			}
			senderReportData, err := senderReport.Marshal()
			if err != nil {
				log.Fatal(err)
			}
			_, err = s.rtcpConn.WriteTo(senderReportData, s.RtcpClientAddr)
			if err != nil {
				s.SenderMutex.RUnlock()
				return
			}
			s.SenderMutex.RUnlock()
			time.Sleep(time.Second)
		}
	}()
}

func (s *Streamer) ParseRtcpPacket(packets []rtcp.Packet) {
	var err error
	for _, packet := range packets {
		rawPacket, _ := packet.Marshal()
		zap.L().Sugar().Infow("", "packet type", reflect.TypeOf(packet))
		switch packet.(type) {
		case *rtcp.CompoundPacket:
			compoundPacket := rtcp.CompoundPacket{}
			err := compoundPacket.Unmarshal(rawPacket)
			if err != nil {
				return
			}
			s.ParseRtcpPacket(compoundPacket)
		case *rtcp.ReceiverReport:
			receiverReport := rtcp.ReceiverReport{}
			err = receiverReport.Unmarshal(rawPacket)
			if err != nil {
				return
			}
			for _, report := range receiverReport.Reports {
				zap.L().Sugar().Infow("Receiver report", "jitter", report.Jitter)
			}
		case *rtcp.Goodbye:
			byePacket := rtcp.Goodbye{}
			err = byePacket.Unmarshal(rawPacket)
			if err != nil {
				return
			}
		}
	}
}
func (s *Streamer) SendPacket(data []byte, pts uint32) error {
	var fragments [][]byte
	if len(data) > MTU {
		fragments = packetizeNalUnit(data, DefaultFragmentationType, MTU)
	} else {
		fragments = [][]byte{data}
	}
	s.SenderMutex.Lock()
	defer s.SenderMutex.Unlock()
	s.rtpTimestamp += s.tsIncrement
	s.lastRtpTimestamp = s.rtpTimestamp
	marker := 0

	s.lastRtpPacket = time.Now()

	for i, f := range fragments {
		if i == len(fragments)-1 {
			marker = 1
		}
		rtpHeader := parser.Header{
			Version:        parser.RtpVersion,
			Marker:         byte(marker),
			PayloadType:    s.PayloadType,
			SequenceNumber: uint16(s.GetNextSeqNumber()),
			Timestamp:      s.rtpTimestamp,
			SSRC:           s.SSRC,
		}
		rtpPacket := parser.SerializeRTPPacket(rtpHeader, f)
		_, err := s.rtpConn.WriteTo(rtpPacket, s.ClientAddr)
		if err != nil {
			zap.L().Sugar().Debugw("failed to send packet", "stream id", s.streamID, "client addr", s.ClientAddr)
			return err
		}
		s.packetCount += 1
	}
	return nil
}
func (s *Streamer) GetNextSeqNumber() int32 {
	if s.SequenceNumber == 0 {
		random := rand.New(rand.New(rand.NewSource(time.Now().Unix())))
		s.SequenceNumber = random.Int31()
	}
	s.SequenceNumber += 1
	return s.SequenceNumber
}
