package main

import "net"

const (
	RRPacket = 201
)

type RtpSession struct {
	RtpIp    net.IPAddr
	RtpPort  int
	RtcpIP   net.IPAddr
	RtcpPort int
}

type RtcpHeader struct {
	Version    byte
	Padding    byte
	ItemCount  byte
	PacketType byte
	Length     uint32
}

type ReceiverReport struct {
	ReporterSSRC       uint32
	ReporteeSSRC       uint32
	LostFraction       uint32
	NumberPacketsLoss  uint32 // this actually a 24 bit number
	HighestSeqNumber   uint32
	InterarrivalJitter uint32
	LSR                uint32 // means: Timestamp of last sender report received
	DLSR               uint32 // means: Delay since last sender report received
}
