package rtsp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

const LineBreak = "\r\n"

const RtspVersion = "RTSP/1.0"

const (
	TransportHeader     = "Transport"
	CSeqHeader          = "CSeq"
	ContentLengthHeader = "Content-Length"
	SessionHeader       = "Session"
)

type Serializable interface {
	Serialize() string
}

type Request struct {
	Method  string
	Uri     *url.URL
	Version string
	Headers map[string]string
	Body    []byte
	session Session
}

type Transport struct {
	Profile        string
	LowerTransport string
	Append         bool
	IsUnicast      bool
	Destination    string
	Interleaved    string
	TTL            string
	Layers         string
	Port           string
	ClientPort     string
	ServerPort     string
	Ssrc           string
	Mode           string
}

type Response struct {
	Version    string
	StatusCode int
	Reason     string
	Headers    map[string]string
	Body       []byte
}

func ParseRequest(reader io.Reader) (Request, error) {
	rtpMsg := Request{
		Headers: map[string]string{
			"Content-Length": "0",
		},
	}
	var i = 0
	bufferedReader := bufio.NewReader(reader)
	currentLine, err := readline(bufferedReader)
	if err != nil {
		return Request{}, err
	}
	fmt.Printf("new index: %v\n", i)
	statusLine := strings.Split(currentLine, " ")
	fmt.Printf("status line: %v\n", statusLine)
	if len(statusLine) != 3 {
		return Request{}, errors.New("too many item in status line")
	}
	rtpMsg.Method = statusLine[0]
	u, err := url.Parse(statusLine[1])
	if err != nil {
		return Request{}, err
	}
	rtpMsg.Uri = u
	rtpMsg.Version = statusLine[2]
	err = parseHeader(bufferedReader, rtpMsg.Headers)
	if err != nil {
		fmt.Printf("incorrect header line")
		return Request{}, err
	}

	// check if for request body

	bodyLen, err := strconv.Atoi(rtpMsg.Headers["Content-Length"])
	if err != nil {
		fmt.Println("Invalid content length")
		return Request{}, err
	}

	rtpMsg.Body = make([]byte, bodyLen)

	_, err = bufferedReader.Read(rtpMsg.Body)
	if err != nil {
		fmt.Println("Failed to read request body")
		return Request{}, err
	}

	return rtpMsg, nil
}
func ParseResponse(reader io.Reader) (Response, error) {
	rtspResponse := Response{
		Headers: map[string]string{},
	}

	bufferedReader := bufio.NewReader(reader)
	currentLine, err := readline(bufferedReader)
	statusLine := strings.Split(currentLine, " ")
	if len(statusLine) != 3 {
		return Response{}, errors.New("too many item in status line")
	}
	rtspResponse.Version = statusLine[0]
	statusCode, err := strconv.Atoi(statusLine[1])
	if err != nil {
		return Response{}, errors.New("incorrect status code")
	}
	rtspResponse.StatusCode = statusCode
	rtspResponse.Reason = statusLine[2]
	err = parseHeader(bufferedReader, rtspResponse.Headers)
	if err != nil {
		fmt.Printf("incorrect header line")
		return Response{}, err
	}

	// check if for request body

	bodyLen, err := strconv.Atoi(rtspResponse.Headers["Content-Length"])
	if err != nil {
		fmt.Println("Invalid content length")
		return Response{}, err
	}

	rtspResponse.Body = make([]byte, bodyLen)

	_, err = bufferedReader.Read(rtspResponse.Body)
	if err != nil {
		fmt.Println("Failed to read request body")
		return Response{}, err
	}

	return rtspResponse, nil
}
func serializeResponse(response Response) string {
	rawResponse := fmt.Sprintf("%d %s %s \r\n", response.StatusCode, response.Reason, response.Version)
	for key, val := range response.Headers {
		rawResponse += fmt.Sprintf("%s: %s\r\n", key, val)
	}
	rawResponse += "\r\n"
	rawResponse += string(response.Body)
	return rawResponse
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
func ParseTransport(input string) []Transport {
	transportsString := strings.Split(input, ",")
	var transports []Transport
	for i, transport := range transportsString {
		t := Transport{}
		transports = append(transports, Transport{})
		components := strings.Split(transport, ";")
		t.Profile = components[0]

		for j := 1; j < len(components); j++ {
			paramComponents := strings.Split(components[j], "=")
			key := strings.TrimSpace(paramComponents[0])
			hasValue := len(paramComponents) >= 2
			switch key {
			case "unicast":
				transports[i].IsUnicast = true
			case "multicast":
				transports[i].IsUnicast = false
			case "destination":
				if hasValue {
					transports[i].Destination = paramComponents[1]
				}
			case "layers":
				transports[i].Layers = paramComponents[1]
			case "interleaved":
				transports[i].Interleaved = paramComponents[1]
			case "port":
				transports[i].Port = paramComponents[1]
			case "client_port":
				transports[i].ClientPort = paramComponents[1]
			case "server_port":
				transports[i].ServerPort = paramComponents[1]
			case "ssrc":
				transports[i].Ssrc = paramComponents[1]
			case "mode":
				transports[i].Mode = paramComponents[1]
			case "ttl":
				transports[i].TTL = paramComponents[1]
			}

		}
	}

	return transports
}

func (t Transport) Serialize() string {
	return ""
}
