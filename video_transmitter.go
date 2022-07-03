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
const addr = ":4874"
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
	listener, err := net.Listen("udp", addr)
	if err != nil {
		fmt.Printf("failed to create udp connection\nerr:%v", err)
	}
	for {

		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go handleRtpSession(context.Background(), conn)

	}

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

func handleRtpSession(ctx context.Context, conn net.Conn) {
	buff := make([]byte, mtu)
	for {
		read, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("failed to read from udp conn\nerr:%v", err)
			return
		}
		_ = read

		fmt.Printf("receive new message")
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
