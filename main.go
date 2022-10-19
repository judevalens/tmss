package main

import (
	"fmt"
	"tmss/rtsp/schemas/request"
)

func main() {

	req := "OPTIONSs rtsp://example.com/1090/920/media.mp4 RTSP/1.0\n" +
		"CSeq: 1\n" +
		"Require: implicit-play\n" +
		"Proxy-Require: gzipped-messages\n" +
		"\n"
	rtspReq, err := request.ParseRtspReq(req)

	if err != nil {
		fmt.Printf("err in main: %v\n", err)
		return
	}

	fmt.Printf("method: %v\n", rtspReq.Request.Method)
}
