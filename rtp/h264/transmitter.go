package h264

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "/usr/local/usr/include/video_reader.h"
import "C"
import (
	"errors"
	"math"
	"net"
	"time"
	"tmss/media"
	"tmss/rtp/parser"
	"tmss/rtsp/headers"
	"unsafe"
)

const MTU = 65000
const NBuffer = 2
const BufferSize = 60
const BuffSizeFactor = 5
const BitsInByte = 8
const (
	GetNextPacket = iota
	StopStream    = iota
	Play          = iota
	Pause         = iota
	Idle          = iota
	Seek
)

type AvPacket *C.struct_AVPakcket
type MediaStreamer interface {
	Play(timeRange headers.Range)
	Pause(timeRange headers.Range)
}

type Buffer struct {
	CurrentBuffer    *C.struct_MediaBuffer
	currentBufferIdx int
	buffChan         chan int
	isBuffering      bool
}

type Streamer struct {
	mediaId string
	media.RepoI
	buffer        Buffer
	bufferCommand chan BufferCommand
	rtpConn       net.PacketConn
	OutByteRate   int
}

func Init(mediaRecord media.Media, streamId int, rtpConn net.PacketConn, rtcpConn net.PacketConn) (*Streamer, error) {
	streamer := &Streamer{}
	streamer.OutByteRate = int(math.Ceil(float64(mediaRecord.Streams[streamId].BitRate) / float64(BitsInByte)))
	buffSizeByte := C.int(BuffSizeFactor * streamer.OutByteRate)
	// TODO change Stream.Path to Stream.Name
	streamPath := C.CString(media.GetStreamPath(mediaRecord.Streams[streamId].Path))
	mediaBuffer := C.init_media_buffer(streamPath, buffSizeByte)
	if mediaBuffer == nil {
		return nil, errors.New("failed to create media buffer")
	}
	return streamer, nil
}

type BufferCommand struct {
	position headers.Range
	command  int
}

func (s Streamer) Play(timeRange headers.Range) {
	//TODO implement me
	panic("implement me")
}

func (s Streamer) Pause(timeRange headers.Range) {
	panic("implement me")
}

func (s Streamer) startServer(rtpConn net.PacketConn, control chan int, streamId int) {
	buff := make([]byte, MTU)
	nBytesRead, a, err := rtpConn.ReadFrom(buff)
	mediaData := s.GetMedia(s.mediaId)
	stream := mediaData.Streams[streamId]
	packet := parser.ParseRtpPacket(buff, nBytesRead)
	_ = packet
	_ = a
	_ = stream
	lastCommand := BufferCommand{
		command: Idle,
	}
	t0 := 0

	maxByteOut := s.OutByteRate
	remainingSize := maxByteOut
	go s.buffer.BufferUp()
	for {
		select {
		case newCommand := <-s.bufferCommand:
			lastCommand = newCommand
		default:
			switch lastCommand.command {
			case Play:
				{
					if t0 <= 0 || time.Now().Sub(time.Unix(0, int64(t0))) > time.Second {
						t0 = time.Now().Nanosecond()
						remainingSize = maxByteOut
					}
					avPacket := s.buffer.ReadNextPacket(true)
					if int(avPacket.size) > remainingSize {
						continue
					}
					avPacket = s.buffer.ReadNextPacket(false)
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
		if err != nil {
			return
		}
	}
}

func (buffer Buffer) ReadNextPacket(peek bool) *C.struct_AVPacket {
	for {
		//TODO need to check pos is within the BOUND
		CurrentBufferPtr := unsafe.Pointer(buffer.CurrentBuffer.packetBuffers)
		currentBuff := *(**C.struct_PacketBuffer)(unsafe.Add(CurrentBufferPtr, unsafe.Sizeof(buffer.CurrentBuffer.packetBuffers)*uintptr(buffer.currentBufferIdx)))
		// the flag isBuffering prevents re-buffering the next queue before it is used or is being buffered
		if float32(currentBuff.currentIdx)/float32(currentBuff.size) <= 0.5 && !buffer.isBuffering {
			buffer.isBuffering = !buffer.isBuffering
			nextBuffIdx := (buffer.currentBufferIdx + 1) % NBuffer
			buffer.buffChan <- nextBuffIdx
		}
		if currentBuff.currentIdx == currentBuff.size {
			buffer.currentBufferIdx = <-buffer.buffChan
			buffer.isBuffering = !buffer.isBuffering
		}
		packets := unsafe.Pointer(currentBuff.packets)
		currentPacket := *(**C.struct_AVPacket)(unsafe.Add(packets, unsafe.Sizeof(currentBuff.packets)*uintptr(currentBuff.currentIdx)))
		if !peek {
			currentBuff.currentIdx = currentBuff.currentIdx + 1
			currentBuff.currentByteSize -= currentPacket.size
		}
		return currentPacket
	}
}

func (buffer Buffer) BufferUp() {
	for {
		bufferIdx := <-buffer.buffChan
		C.buffer(buffer.CurrentBuffer, C.int(bufferIdx))
		buffer.buffChan <- bufferIdx
	}
}
