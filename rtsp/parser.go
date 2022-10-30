package rtsp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
)

const (
	ANNOUNCE     = "ANNOUNCE"
	DESCRIBE     = "DESCRIBE"
	GetParameter = "GET_PARAMETER"
	OPTIONS      = "OPTIONS"
	PAUSE        = "PAUSE"
	PLAY         = "PLAY"
	RECORD       = "RECORD"
	REDIRECT     = "REDIRECT"
	SETUP        = "SETUP"
	SetParameter = "SET_PARAMETER"
	TEARDOWN     = "TEARDOWN"
)

const LineBreak = "\r\n"

const RtspVersion = "RTSP/1.0"
const (
	TimeLayout = "15:04:05.000"
)
const (
	TransportHeader     = "Transport"
	CSeqHeader          = "Cseq"
	ContentLengthHeader = "Content-Length"
	SessionHeader       = "Session"
	RangeHeader         = "Range"
)

type Serializable interface {
	Serialize() string
}

func ParseRequest(reader io.Reader) (*http.Request, error) {
	req := &http.Request{}
	var i = 0
	bufferedReader := bufio.NewReader(reader)
	currentLine, err := readline(bufferedReader)
	if err != nil {
		return nil, err
	}
	fmt.Printf("new index: %v\n", i)
	statusLine := strings.Split(currentLine, " ")
	fmt.Printf("status line: %v\n", statusLine)
	if len(statusLine) != 3 {
		return nil, errors.New("too many item in status line")
	}
	req.Method = statusLine[0]
	reqUrl, err := url.Parse(statusLine[1])
	if err != nil {
		return nil, err
	}
	req.URL = reqUrl
	req.Proto = statusLine[2]

	headerReader := textproto.NewReader(bufferedReader)
	header, err := headerReader.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}
	req.Header = http.Header(header)
	_, found := req.Header[ContentLengthHeader]
	if !found {
		req.Header.Add(ContentLengthHeader, "0")
	}
	bodyLen, err := strconv.Atoi(req.Header.Get(ContentLengthHeader))
	if err != nil {
		return nil, err
	}
	bodyBuff := make([]byte, bodyLen)
	_, err = bufferedReader.Read(bodyBuff)
	if err != nil {
		return nil, err
	}
	// TODO should user a proper Closer
	req.Body = io.NopCloser(bytes.NewReader(bodyBuff))
	if err != nil {
		fmt.Println("Failed to read request body")
		return nil, err
	}

	return req, nil
}
func ParseResponse(reader io.Reader) (*http.Response, error) {
	response := &http.Response{}

	bufferedReader := bufio.NewReader(reader)
	currentLine, err := readline(bufferedReader)
	proto, status, found := strings.Cut(currentLine, " ")
	if !found {
		return nil, errors.New("invalid status line")
	}
	response.Proto = proto

	statusCode, _, found := strings.Cut(status, " ")

	if !found {
		return nil, errors.New("invalid status line")
	}

	response.StatusCode, err = strconv.Atoi(statusCode)
	if err != nil {
		return nil, errors.New("incorrect status code")
	}
	response.Status = status

	headerReader := textproto.NewReader(bufferedReader)
	header, err := headerReader.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}
	response.Header = http.Header(header)
	_, found = response.Header[ContentLengthHeader]
	if !found {
		response.Header.Add(ContentLengthHeader, "0")
	}
	bodyLen, err := strconv.Atoi(response.Header.Get(ContentLengthHeader))
	if err != nil {
		return nil, err
	}
	bodyBuff := make([]byte, bodyLen)
	_, err = bufferedReader.Read(bodyBuff)
	if err != nil {
		return nil, err
	}
	// TODO should user a proper Closer
	response.Body = io.NopCloser(bytes.NewReader(bodyBuff))
	if err != nil {
		fmt.Println("Failed to read request body")
		return nil, err
	}
	return response, nil
}
func serializeResponse(response http.Response) string {
	rawResponse := fmt.Sprintf("%s %d %s \r\n", response.Proto, response.StatusCode, response.Status)
	for key, val := range response.Header {
		rawResponse += fmt.Sprintf("%s: %s\r\n", key, val)
	}
	rawResponse += "\r\n"
	buff := make([]byte, response.ContentLength)
	_, err := response.Body.Read(buff)
	if err != nil {
		return ""
	}
	rawResponse += string(buff)
	return rawResponse
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
