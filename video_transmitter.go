package main

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "lib/mms/video_reader.h"
import "C"
import (
	"context"
	"fmt"
	"net"
)

const receiverAdd = ""
const addr = "[::1]:5009"
const mtu = 1200

const MaxFrameBufferSize = 10

type VideoTransmitter struct {
	mediaBuffer *C.struct_MediaBuffer
	packetChan  chan chan NalUnit
}

func initVideoTransmitter() *VideoTransmitter {
	videoTransmitter := &VideoTransmitter{}
	videoTransmitter.mediaBuffer = C.init_media_buffer(C.CString("file:/mnt/c/Users/judev/Downloads/big buck bunny.mp4"))
	C.bufferUP(videoTransmitter.mediaBuffer)
	return videoTransmitter
}

func startRtpServer() {

	var err error
	var buffer = make([]byte, 25000)
	listener, err := net.ListenPacket("udp6", addr)
	if err != nil {
		fmt.Printf("failed to create udp connection\nerr:%v\n", err)
		return
	}
	fmt.Printf("waiting for packets on %v\n", listener.LocalAddr())
	for {
		_, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			return
		}

		handleRtpSession(context.Background(), buffer, addr)

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

func handleRtpSession(ctx context.Context, msg []byte, conn net.Addr) {
	fmt.Printf("receive new message\n")
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

func (t VideoTransmitter) buildHdrHeader() NAlHeader {
	return NAlHeader{}
}

func FillBuffer() {

}
