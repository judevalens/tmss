package rtsp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
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
	Body    []byte
}

type Transport struct {
	Append      bool
	IsUnicast   bool
	Destination net.Addr
	Interleaved string
	TTL         int
	Layers      int
	Port        string
	ClientPort  string
	ServerPort  string
	Ssrc        string
	Mode        string
}

type RtspResponse struct {
	Version    string
	StatusCode int
	Reason     string
	Headers    map[string]string
}

func ParseRequest(reader io.Reader) (RtspRequest, error) {
	rtpMsg := RtspRequest{
		Headers: map[string]string{
			"Content-Length": "0",
		},
	}
	var i = 0
	bufferedReader := bufio.NewReader(reader)
	currentLine, err := readline(bufferedReader)
	if err != nil {
		return RtspRequest{}, err
	}
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
	err = parseHeader(bufferedReader, rtpMsg.Headers)
	if err != nil {
		fmt.Printf("incorrect header line")
		return RtspRequest{}, err
	}

	// check if for request body

	bodyLen, err := strconv.Atoi(rtpMsg.Headers["Content-Length"])
	if err != nil {
		fmt.Println("Invalid content length")
		return RtspRequest{}, err
	}

	rtpMsg.Body = make([]byte, bodyLen)

	_, err = bufferedReader.Read(rtpMsg.Body)
	if err != nil {
		fmt.Println("Failed to read request body")
		return RtspRequest{}, err
	}

	return rtpMsg, nil
}

func ParseResponse(reader io.Reader) (RtspResponse, error) {
	rtspResponse := RtspResponse{
		Headers: map[string]string{},
	}

	bufferedReader := bufio.NewReader(reader)
	currentLine, err := readline(bufferedReader)
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
	err = parseHeader(bufferedReader, rtspResponse.Headers)
	if err != nil {
		fmt.Printf("incorrect header line")
		return RtspResponse{}, err
	}
	return rtspResponse, nil
}

func parseHeader(reader *bufio.Reader, headers map[string]string) error {
	currentLine, err := readline(reader)
	for currentLine != "" && err == nil {
		key, value, found := strings.Cut(currentLine, ":")
		/// check that grammar is valid
		if !found && currentLine != LineBreak {
			fmt.Printf("incorrect header line: %v\n", currentLine)
			return errors.New("incorrect header line")
		}
		if found {
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)
			headers[key] = value
		}
		currentLine, err = readline(reader)
	}
	return nil
}

func readline(reader *bufio.Reader) (string, error) {
	var buffer string
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			buffer += readString
			return buffer, err
		}
		buffer += readString

		if buffer[len(buffer)-2:] == "\r\n" {
			return buffer[:len(buffer)-2], nil
		}
	}

}
