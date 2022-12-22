package parser

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

const RtpHeaderSize = 16
const RtpVersion = 2

const (
	H264PayloadType = iota
	AccPayloadType  = iota
)

type Packet struct {
	Header  Header
	Payload []byte
}

type Header struct {
	Version             byte
	Padding             byte
	Extension           byte
	CsrcCount           byte
	Marker              byte
	PayloadType         byte
	SequenceNumber      uint16
	Timestamp           uint32
	SSRC                uint32
	CSRC                []uint32
	ExtensionHeaderId   uint16
	ExtensionHeaderSize uint16
	ExtensionHeader     []uint32
	Size                int
}

type RtpPayload interface {
}

func parseRtpHeader(data []byte) Header {
	fmt.Printf("dec = %v, binary str =  %v\n", data[0], strconv.FormatInt(int64(data[0]), 2))
	header := Header{
		Version:        (data[0] << 6) >> 6,
		Padding:        (data[0] << 5) >> 7,
		Extension:      (data[0] << 4) >> 7,
		CsrcCount:      data[0] >> 4,
		Marker:         (data[1] << 7) >> 7,
		PayloadType:    (data[1] << 1) >> 1,
		SequenceNumber: binary.BigEndian.Uint16(data[2:4]),
		Timestamp:      binary.BigEndian.Uint32(data[4:8]),
		SSRC:           binary.BigEndian.Uint32(data[8:12]),
	}
	i := 12
	for ; i < (int)(header.CsrcCount); i += 4 {
		header.CSRC = append(header.CSRC, binary.BigEndian.Uint32(data[i:i+4]))
	}
	if header.Extension == 1 {
		i = int(12 + 4*uint(header.CsrcCount))
		j := i + 2
		header.ExtensionHeaderId = binary.BigEndian.Uint16(data[i:j])
		i = j + 2
		j = i + 2
		header.ExtensionHeaderSize = binary.BigEndian.Uint16(data[i:j])
		for i = j; i < i+int(header.ExtensionHeaderSize); i += 4 {
			header.ExtensionHeader = append(header.ExtensionHeader, binary.BigEndian.Uint32(data[i:i+4]))
		}
	}
	header.Size = i // will use the header size to determine the payload size
	return header
}

func serializeRtpHeader(header Header) []byte {

	rawHeader := make([]byte, RtpHeaderSize+(4*(int)(header.CsrcCount)))
	rawHeader[0] = rawHeader[0] | (header.Padding << 2)
	rawHeader[0] = rawHeader[0] | (header.Extension << 3)
	rawHeader[0] = rawHeader[0] | (header.CsrcCount << 4)
	rawHeader[1] = rawHeader[1] | (header.Marker)
	rawHeader[1] = rawHeader[1] | (header.PayloadType << 1)
	binary.BigEndian.PutUint32(rawHeader[2:4], uint32(header.CsrcCount))
	binary.BigEndian.PutUint32(rawHeader[4:8], header.Timestamp)
	binary.BigEndian.PutUint32(rawHeader[8:12], header.SSRC)

	for i := 0; i < (int)(header.CsrcCount); i++ {
		s := 12 + (4 * i)
		binary.BigEndian.PutUint32(rawHeader[s:s+4], header.CSRC[i])
	}
	return rawHeader
}

func SerializeRTPPacket(header []byte, payload []byte) []byte {
	rawPacket := make([]byte, len(header)+len(payload))
	copy(rawPacket, header)
	copy(rawPacket[len(header):], payload)
	return rawPacket
}

func ParseRtpPacket(packet []byte, packetSize int) Packet {
	rtpHeader := parseRtpHeader(packet)
	return Packet{
		Header:  rtpHeader,
		Payload: packet[rtpHeader.Size:packetSize],
	}
}
func getRtpSequence() int16 {
	return 1
}
