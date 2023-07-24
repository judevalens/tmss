package h264

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "/usr/local/usr/include/video_reader.h"
import "C"
import (
	"go.uber.org/zap"
	"unsafe"
)

const MTU = 65000
const NBuffer = 2
const BufferSize = 60
const BuffSizeFactor = 5
const BitsInByte = 8

type Buffer struct {
	CurrentBuffer    *C.struct_MediaBuffer
	currentBufferIdx int
	buffChan         chan int
	isBuffering      bool
	init             bool
}

func (buffer *Buffer) ReadNextPacket(peek bool) *C.struct_AVPacket {
	for {
		//TODO need to check pos is within the BOUND
		CurrentBufferPtr := unsafe.Pointer(buffer.CurrentBuffer.packetBuffers)
		currentBuff := *(**C.struct_PacketBuffer)(unsafe.Add(CurrentBufferPtr, unsafe.Sizeof(buffer.CurrentBuffer.packetBuffers)*uintptr(buffer.currentBufferIdx)))
		zap.L().Sugar().Debugw("reading new packet", "id", currentBuff.currentIdx)
		// the flag isBuffering prevents re-buffering the next queue before it is used or is being buffered
		if (float32(currentBuff.currentIdx)/float32(currentBuff.size) >= 0.5) || !buffer.init {
			if !buffer.init {
				buffer.init = true
				buffer.buffChan <- 0
				buffer.currentBufferIdx = <-buffer.buffChan
				currentBuff = *(**C.struct_PacketBuffer)(unsafe.Add(CurrentBufferPtr, unsafe.Sizeof(buffer.CurrentBuffer.packetBuffers)*uintptr(buffer.currentBufferIdx)))
			} else {
				if !buffer.isBuffering {
					buffer.isBuffering = !buffer.isBuffering
					nextBuffIdx := (buffer.currentBufferIdx + 1) % NBuffer
					buffer.buffChan <- nextBuffIdx
				}
				// checks if current buffer is empty
				if currentBuff.currentIdx == currentBuff.size {
					if currentBuff.eof == 1 {
						zap.L().Info("end of stream")
						return nil
					}
					buffer.currentBufferIdx = <-buffer.buffChan
					currentBuff = *(**C.struct_PacketBuffer)(unsafe.Add(CurrentBufferPtr, unsafe.Sizeof(buffer.CurrentBuffer.packetBuffers)*uintptr(buffer.currentBufferIdx)))
					currentBuff.currentIdx = 0
					buffer.isBuffering = !buffer.isBuffering
				}
			}
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

func (buffer *Buffer) BufferUp() {
	for {
		bufferIdx := <-buffer.buffChan
		C.buffer(buffer.CurrentBuffer, C.int(bufferIdx))
		buffer.buffChan <- bufferIdx
	}
}
