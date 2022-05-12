package main

type RtpHeader struct {
	Version        byte
	Padding        byte
	Extension      byte
	CrcCount       byte
	Marker         byte
	PayloadType    byte
	SequenceNumber int16
	Timestamp      int32
	SSRC           int32
	CSRC           int32
}
