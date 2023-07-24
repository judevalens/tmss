package h264

import (
	"math"
)

type FUA struct {
	FuIndicator NAlHeader
	FuHeader    FuHeader
	DoN         int16
	Payload     []byte
}

// packetizes a NAL unit according to RFC 6184
//
// The packet will be fragmented into multiple rtp payload if the payload is larger than the MTU
func packetizeNalUnit(payload []byte, fragmentationType byte, maxPacketSize int) [][]byte {
	nFragment := (int)(math.Ceil(float64(len(payload)) / float64(maxPacketSize)))
	rawPackets := make([][]byte, nFragment)
	var fuIndicator byte
	singleNalHeader := payload[0]
	nri := (singleNalHeader << 1) >> 6
	nalType := (singleNalHeader << 3) >> 3
	//TODO add forbidden bit
	fuIndicator = fuIndicator | (nri << 5)
	fuIndicator = fuIndicator | (fragmentationType)
	startIndex := 0
	endIndex := 0
	for i := 0; i < nFragment; i++ {
		var fuHeader byte
		if i == 0 {
			fuHeader = fuHeader | (0x01 << 7)
		} else if i == nFragment-1 {
			fuHeader = fuHeader | (0x01 << 6)
		}
		fuHeader = fuHeader | (nalType)
		endIndex = startIndex + int(math.Min(float64(maxPacketSize), float64(len(payload)-startIndex)))
		rawPackets[i] = make([]byte, 2+endIndex-startIndex)
		rawPackets[i][0] = fuIndicator
		rawPackets[i][1] = fuHeader
		copy(rawPackets[i][2:], payload[startIndex:endIndex])
		startIndex = endIndex
	}
	return rawPackets
}
