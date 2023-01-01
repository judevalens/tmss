package h264

import "encoding/binary"

type Header struct {
	Version       byte
	Profile       byte
	Compatibility byte
	Level         byte
	Reserved      byte
	NLULengthSize byte
	Reserved2     byte
	SPSCount      byte
	SPSSize       uint16
	SPS           []byte
	PPSCount      byte
	PPSSize       uint16
	PPS           []byte
}

func ParseAvccHeader(data []byte) Header {
	// ref : https://stackoverflow.com/questions/24884827/possible-locations-for-sequence-picture-parameter-sets-for-h-264-stream
	header := Header{
		Version:       data[0],
		Profile:       data[1],
		Compatibility: data[2],
		Level:         data[3],
		NLULengthSize: (data[4] >> 6) + 1,
		SPSCount:      data[5] >> 3,
	}

	header.SPSSize = binary.BigEndian.Uint16(data[6:8])
	i := 8
	e := i + int(header.SPSSize)
	header.SPS = data[i:e]
	i = e
	e += 1
	header.PPSCount = data[i]
	i = e
	e += 2
	header.PPSSize = binary.BigEndian.Uint16(data[i:e])
	i = e
	e += int(header.PPSSize)
	header.PPS = data[i:e]
	return header
}

// AVCCToNalU parses extract NAL units from a block of data encoded in AVCC
func AVCCToNalU(header Header, payload []byte) [][]byte {
	nalUnits := make([][]byte, 0)
	i := 0
	for i < len(payload) {
		var nalSize uint32
		end := i
		switch header.NLULengthSize {
		case 1:
			end += 1
			nalSize = uint32(payload[i:end][0])
		case 2:
			end += 2
			nalSize = uint32(binary.BigEndian.Uint16(payload[i:end]))
		case 4:
			end += 4
			nalSize = binary.BigEndian.Uint32(payload[i:end])
		}
		i = end
		end += int(nalSize)
		nalUnits = append(nalUnits, payload[i:end])
		i = end
	}

	return nalUnits
}
