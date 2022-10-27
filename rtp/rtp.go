package rtp

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

type RtpPacket struct {
	Header  RtpHeader
	payload RtpPayload
}

type RtpHeader struct {
	Version        byte
	Padding        byte
	Extension      byte
	CsrcCount      byte
	Marker         byte
	PayloadType    byte
	SequenceNumber uint16
	Timestamp      uint32
	SSRC           uint32
	CSRC           []uint32
}

type RtpPayload interface {
}

func parseRtpHeader(data []byte) RtpHeader {

	println()
	fmt.Printf("dec = %v, binary str =  %v\n", data[0], strconv.FormatInt(int64(data[0]), 2))
	header := RtpHeader{
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
	/**for i := 0; i < (int)(headers.CsrcCount); i++ {
			s := 12 + (4 * i)
			headers.CSRC = append(headers.CSRC, binary.BigEndian.Uint32(data[s:s+4]))
		}

		fmt.Printf("%v", data[:16])
	**/
	return header
}

func serializeRtpHeader(header RtpHeader) []byte {

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

func ParseRtpPacket(packet []byte, packetSize int) RtpPacket {
	rtpHeader := parseRtpHeader(packet)
	var rtpPayload RtpPayload
	switch rtpHeader.PayloadType {
	case H264PayloadType:
		rtpPayload = parseH264Payload()
	case AccPayloadType:
		rtpPayload = parseAccPayload()
	}
	return RtpPacket{
		Header:  rtpHeader,
		payload: rtpPayload,
	}
}
func getRtpSequence() int16 {
	return 1
}

func parseH264Payload() RtpPayload {
	//TODO	log.Fatal("Need to be implemented")
	return SingleNalPacket{}
}
func parseAccPayload() RtpPayload {
	//TODO wlog.Fatal("Need to be implemented")
	return SingleNalPacket{}
}
