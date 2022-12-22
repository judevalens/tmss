package h264

import (
	"fmt"
	"math"
	"strconv"
	"tmss/rtp/parser"
)

const (
	NALU   = 23
	StapA  = 24
	StapB  = 25
	Mtap16 = 26
	Mtap24 = 27
	FuA    = 28
	FuB    = 29
)

type NAlHeader struct {
	ForbiddenBit byte
	NRI          byte
	NalType      byte
}

type FuHeader struct {
	Start    byte
	End      byte
	Reserved byte
	Type     byte
}

type FragmentationUnitHeader struct {
	StartBit    byte
	EndByte     byte
	ReservedBit byte
	Type        byte
}

type H264Payload struct {
	header NAlHeader
	data   []byte
}

type NalUnit interface {
	serialize() []byte
}

type SingleNalPacket struct {
	header NAlHeader
	data   []byte
}

type StapAunit struct {
	header NAlHeader
	data   []byte
}

type StapBunit struct {
	header NAlHeader
	DON    int16
	data   []byte
}

type Mtap16Unit struct {
	Header NAlHeader
	DOND   int8
	TsOffset int16
	Payload  []byte
}
type Mtap24Unit struct {
	Header NAlHeader
	DOND   int8
	TsOffset int32
	Payload  []byte
}
type StapApacketContainer struct {
	header  NAlHeader
	packets []StapAunit
}
type NalAggregationPacket interface {
	StapApacketContainer | StapBpacketContainer
}
type NalAggregationPacketContainer interface {
	packetize(data []byte)
}

func (p *StapApacketContainer) getNext() SingleNalPacket {
	panic("implement me")
}
func (p *StapApacketContainer) packetize(data []byte) {
	panic("implement me")
}

type StapBpacketContainer struct {
	DON     int8
	header  NAlHeader
	packets []StapBunit
}

type Mtap16PacketContainer struct {
	header NAlHeader
	DONB   int16
	Units  []Mtap24Unit
}
type Mtap24PacketContainer struct {
	header NAlHeader
	DONB   int16
	Units  []Mtap24Unit
}

type FUa struct {
	FuIndicator NAlHeader
	FuHeader    FuHeader
	DoN         int16
	Payload     []byte
}

func (packet SingleNalPacket) serialize() []byte {
	var header byte
	rawPacket := make([]byte, len(packet.data)+1)

	rawPacket[0] = header
	copy(rawPacket[1:], packet.data)
	header = header | packet.header.ForbiddenBit
	header = header | (packet.header.NRI << 1)
	header = header | (packet.header.NalType << 3)
	return rawPacket
}

func getNalHeader(header byte) NAlHeader {

	//TODO clean up this mess
	forbiddenByte := header & 1
	fmt.Printf("start: %v\n", strconv.FormatInt(int64(header), 2))
	nri := (header << 5) >> 6
	nalType := header >> 3
	fmt.Printf("end: %v\n", strconv.FormatInt(int64(nri), 2))
	fmt.Printf("forbidden byte: %v, nri: %v, nalType: %v\n", forbiddenByte, nri, nalType)
	//ENDofTODO//
	return NAlHeader{
		ForbiddenBit: (header << 7) >> 7,
		NRI:          (header << 5) >> 6,
		NalType:      header >> 3,
	}
}

func getFuHeader(header byte) FuHeader {
	return FuHeader{
		Start:    (header << 7) >> 7,
		End:      (header << 6) >> 7,
		Reserved: (header << 5) >> 7,
		Type:     header >> 3,
	}
}

func (packet FUa) serialize(header []byte, data []byte, maxPacketSize int) [][]byte {
	nFragment := (int)(math.Ceil(float64(len(data)) / float64(maxPacketSize)))
	rawPackets := make([][]byte, nFragment)
	var fuIndicator byte
	var fuHeader byte

	fuIndicator = fuIndicator | packet.FuIndicator.ForbiddenBit
	fuIndicator = fuIndicator | (packet.FuIndicator.NRI << 2)
	fuIndicator = fuIndicator | (packet.FuIndicator.NalType << 3)

	fuHeader = fuHeader | packet.FuHeader.Start
	fuHeader = fuHeader | (packet.FuHeader.End << 1)
	fuHeader = fuHeader | (packet.FuHeader.Reserved << 2)
	fuHeader = fuHeader | (packet.FuHeader.Type << 3)
	startIndex := 0
	endIndex := 0
	println(nFragment)
	for i := 0; i < nFragment; i++ {
		endIndex = startIndex + int(math.Min(float64(maxPacketSize), float64(len(data)-startIndex)))
		rawPackets[i] = data[startIndex:endIndex]
		payloadFragment := make([]byte, 2+endIndex-startIndex)
		payloadFragment[0] = fuIndicator
		payloadFragment[1] = fuHeader
		copy(payloadFragment[2:], data[startIndex:endIndex])
		rawPackets[i] = parser.SerializeRTPPacket(header, payloadFragment)
		startIndex += maxPacketSize
	}
	return rawPackets
}
