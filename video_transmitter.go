package main

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "lib/mms/video_reader.h"
import "C"
import (
	"context"
	"fmt"
	"net"
	"os"
	"path"
	"tmss/rtp"
)

const receiverAdd = ""
const addr = ":5004"
const mtu = 1200

const MaxFrameBufferSize = 10

type VideoTransmitter struct {
	mediaBuffer *C.struct_MediaBuffer
	packetChan  chan chan rtp.NalUnit
}

func initVideoTransmitter() *VideoTransmitter {
	videoTransmitter := &VideoTransmitter{}
	filepath := "/home/jude/Desktop/amnis server/big_buck_bunny.mp4"
	homeDir, _ := os.UserHomeDir()

	outputDir := path.Join(homeDir, "Desktop/amnis server")

	mediaAvFormatCtx := C.open_media(C.CString(filepath))
	C.demux_file(mediaAvFormatCtx, C.CString(outputDir))
	//C.bufferUP(videoTransmitter.mediaBuffer, outputDir)
	return videoTransmitter
}

func startRtpServer(ctx context.Context, addr2 string) {
	var err error
	var buffer = make([]byte, 25000)
	listener, err := net.ListenPacket("udp", addr2)
	if err != nil {
		fmt.Printf("failed to create udp connection\nerr:%v\n", err)
		return
	}
	fmt.Printf("waiting for packets on %v\n", listener.LocalAddr())
	i := 0
	rtpPacketChannel := make(chan rtp.RtpPacket, 10)
	startPacketHandler(ctx, rtpPacketChannel, 10)
	for {
		packetSize, _, err := listener.ReadFrom(buffer)
		if err != nil {
			return
		}
		i++
		fmt.Printf("#%v: new msg, msg len: %v\n", i, packetSize)
		//	handleRtpSession(context.Background(), buffer, addr)
		packet := rtp.ParseRtpPacket(buffer, packetSize)

		rtpPacketChannel <- packet
	}
	fmt.Printf("server shut down\n")

}

func test() {
	var mediaBuffer *C.struct_MediaBuffer
	mediaBuffer = C.init_media_buffer(C.CString("file:/mnt/c/Users/judev/Downloads/big buck bunny.mp4"))
	C.bufferUP(mediaBuffer)
}

type h264FrameBuffer struct {
	videoFrameQueue [MaxFrameBufferSize]h264Frame
	start           int
	end             int
	currentSize     int
}

func (buff *h264FrameBuffer) add(frame h264Frame) {
	buff.videoFrameQueue[buff.end] = frame
	buff.end++
	buff.end %= MaxFrameBufferSize
	buff.currentSize++
}

func (buff *h264FrameBuffer) get() h264Frame {
	frame := buff.videoFrameQueue[buff.start]
	buff.start++
	buff.start %= MaxFrameBufferSize
	buff.currentSize--
	return frame
}

type h264Frame struct {
	data []byte
	dts  int64
	pts  int64
}

func startPacketHandler(ctx context.Context, rtpPacketChannel chan rtp.RtpPacket, poolSize int) {
	i := 0
	for i < poolSize {
		i++
		childCtx, _ := context.WithCancel(ctx)

		go func(ctx context.Context, rtpChannel2 chan rtp.RtpPacket, workerN int) {
			for {
				select {
				case <-childCtx.Done():
					return
				case rtpPacket := <-rtpChannel2:
					fmt.Printf("worker #%v: new packet received,rtp version = %v, payloadtype = %v,  timestamp = %v, seqNumber = %v\n", workerN, rtpPacket.Header.Version, rtpPacket.Header.PayloadType, rtpPacket.Header.Timestamp, rtpPacket.Header.SequenceNumber)
				}
			}
		}(childCtx, rtpPacketChannel, i)
	}
}

func rtpClient() {
	dial, err := net.Dial("udp6", "[::1]:5009")
	if err != nil {
		fmt.Printf("failed to dial rtp server\nerr:%v\n", err)
		return
	}
	buffer := make([]byte, 1500)
	for {
		fmt.Printf("dialed %v\n", dial.LocalAddr().String())

		_, err := dial.Write([]byte("hello from client"))
		if err != nil {
			fmt.Printf("failed to write to server\n")
			return
		}

		read, err := dial.Read(buffer)
		if err != nil {
			fmt.Printf("failed to receive stream from client...\n err: %v\n", err)
			return
		}

		fmt.Printf("received new msg from rtp server.... msg len: %v\n\n", read)
	}

}

func (t VideoTransmitter) transmitH264() {

	for {
		select {}
	}
}

func sendFrame(conn net.Conn) {
}

func (t VideoTransmitter) bufferUP() {
	for {
		select {}
	}
}

func (t VideoTransmitter) buildHdrHeader() rtp.NAlHeader {
	return rtp.NAlHeader{}
}

func FillBuffer() {

}
