package main

import (
	"fmt"
	"strconv"
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
	Header   NAlHeader
	DOND     int8
	TsOffset int16
	Payload  []byte
}
type Mtap24Unit struct {
	Header   NAlHeader
	DOND     int8
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
	FuHeader    FragmentationUnitHeader
	DoN         int16
	Payload     []byte
}

func getNalHeader(headerB byte) NAlHeader {

	//TODO clean up this mess/
	forbiddenByte := headerB & 1
	fmt.Printf("start: %v\n", strconv.FormatInt(int64(headerB), 2))
	nri := (headerB << 5) >> 6
	nalType := headerB >> 3
	fmt.Printf("end: %v\n", strconv.FormatInt(int64(nri), 2))
	fmt.Printf("forbidden byte: %v, nri: %v, nalType: %v\n", forbiddenByte, nri, nalType)
	//ENDofTODO//
	return NAlHeader{
		ForbiddenBit: headerB & 1,
		NRI:          (headerB << 5) >> 6,
		NalType:      headerB >> 3,
	}
}
