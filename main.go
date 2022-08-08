package main

import (
	"fmt"
	"tmss/rtsp_parser"
)

func main() {

	rtpRequest := "" +
		"SETUP rtsp://example.com/media.mp4/streamid=0 RTSP/10\r\n" +
		"CSeq: 3\r\n" +
		"Transport: RTP/AVP;unicast;client_port=8000-8001\r\n" +
		"\r\n"

	req, err := rtsp_parser.ParseRequest(rtpRequest)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("req: %v\n", req)

}
