package rtsp_parser

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const LineBreak = "\r\n"

type RtspRequest struct {
	Method  string
	Uri     *url.URL
	Version string
	Headers map[string]string
}

type RtspResponse struct {
	Version    string
	StatusCode int
	Reason     string
	Headers    map[string]string
}

func ParseRequest(msg string) (RtspRequest, error) {
	rtpMsg := RtspRequest{
		Headers: map[string]string{
			"Content-Length": "0",
		},
	}
	var i = 0
	currentLine, i := readline(msg, i)

	fmt.Printf("new index: %v\n", i)
	statusLine := strings.Split(currentLine, " ")
	fmt.Printf("status line: %v\n", statusLine)
	if len(statusLine) != 3 {
		return RtspRequest{}, errors.New("too many item in status line")
	}
	rtpMsg.Method = statusLine[0]
	u, err := url.Parse(statusLine[1])
	if err != nil {
		return RtspRequest{}, err
	}
	rtpMsg.Uri = u
	rtpMsg.Version = statusLine[2]
	err = parseHeader(msg, i, rtpMsg.Headers)
	if err != nil {
		fmt.Printf("incorrect header line")
		return RtspRequest{}, err
	}
	return rtpMsg, nil
}

func ParseResponse(msg string) (RtspResponse, error) {
	rtspResponse := RtspResponse{
		Headers: map[string]string{},
	}
	i := 0
	currentLine, i := readline(msg, i)
	statusLine := strings.Split(currentLine, " ")
	if len(statusLine) != 3 {
		return RtspResponse{}, errors.New("too many item in status line")
	}
	rtspResponse.Version = statusLine[0]
	statusCode, err := strconv.Atoi(statusLine[1])
	if err != nil {
		return RtspResponse{}, errors.New("incorrect status code")
	}
	rtspResponse.StatusCode = statusCode
	rtspResponse.Reason = statusLine[2]
	err = parseHeader(msg, i, rtspResponse.Headers)
	if err != nil {
		fmt.Printf("incorrect header line")
		return RtspResponse{}, err
	}
	return rtspResponse, nil
}

func parseHeader(msg string, startIndex int, headerMap map[string]string) error {
	for currentLine, startIndex := readline(msg, startIndex); currentLine != "\r\n"; currentLine, startIndex = readline(msg, startIndex) {
		key, value, found := strings.Cut(currentLine, ":")
		/// check that grammar is valid
		if !found && currentLine != LineBreak {
			fmt.Printf("incorrect header line: %v\n", currentLine)
			return errors.New("incorrect header line")
		}
		if found {
			headerMap[key] = value
		}

	}
	return nil
}

func readline(msg string, i int) (string, int) {
	i2 := strings.Index(msg[i:], LineBreak)
	if i2 < 0 {
		return msg, i2
	}
	absEndOfLine := i + i2
	//line is crlf
	if absEndOfLine == i {
		return msg[i : absEndOfLine+len(LineBreak)], absEndOfLine + len(LineBreak)
	}
	fmt.Printf("new line: from %v->%v, %v\n", i, absEndOfLine, msg[i:absEndOfLine])
	fmt.Printf("EOL\n")
	return msg[i:absEndOfLine], absEndOfLine + len(LineBreak)
}
