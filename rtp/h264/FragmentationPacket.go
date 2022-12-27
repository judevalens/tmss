package h264

import (
	"log"
	"math"
)

type FUA struct {
	FuIndicator NAlHeader
	FuHeader    FuHeader
	DoN         int16
	Payload     []byte
}

func SplitSingleNal(data []byte, fragmentationType byte, maxPacketSize int) [][]byte {
	nFragment := (int)(math.Ceil(float64(len(data)) / float64(maxPacketSize)))
	rawPackets := make([][]byte, nFragment)
	var fuIndicator byte
	singleNalHeader := data[0]

	if singleNalHeader != 0 {
		log.Fatal(singleNalHeader)
	}
	nri := (singleNalHeader << 5) >> 6
	nalType := singleNalHeader >> 3
	fuIndicator = fuIndicator | (nri << 1)
	fuIndicator = fuIndicator | (fragmentationType << 3)
	startIndex := 0
	endIndex := 0
	for i := 0; i < nFragment; i++ {
		var fuHeader byte
		if i == 0 {
			fuHeader = fuHeader | 0x01
		} else if i == nFragment-1 {
			fuHeader = fuHeader | (0x01 << 1)
		}
		fuHeader = fuHeader | (nalType << 3)
		endIndex = startIndex + int(math.Min(float64(maxPacketSize), float64(len(data)-startIndex)))
		println(2+endIndex-startIndex)
		rawPackets[i] = make([]byte, 2+endIndex-startIndex)
		rawPackets[i][0] = fuIndicator
		rawPackets[i][1] = fuHeader
		copy(rawPackets[i][2:], data[startIndex:endIndex])
		startIndex = endIndex
	}
	return rawPackets
}
