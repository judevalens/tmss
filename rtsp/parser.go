package rtsp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const LineBreak = "\r\n"

const RtspVersion = "RTSP/1.0"
const (
	TimeLayout = "15:04:05.000"
)
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

type Range struct {
	startTime float64
	endTime   float64
	liveEvent bool
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

func ParseRange(input string) (Range, error) {
	myRange := Range{}

	// https://www.rfc-editor.org/rfc/rfc2326.html#page-17
	nptHhmmss := regexp.MustCompile("^(?P<timeIso>\\d+:\\d{1,2}:\\d{1,2}(\\.\\d*)*)$")
	nptSec := regexp.MustCompile("^(?P<timeUnix>\\d+(\\.\\d+)*)$")
	nptTime := regexp.MustCompile(fmt.Sprintf("((?P<now>now)|%s|%s)", regexToString(nptHhmmss), regexToString(nptSec)))
	start := regexp.MustCompile(fmt.Sprintf("(?P<start>%s)", nptTime))
	end := regexp.MustCompile(fmt.Sprintf("(?P<end>%s)", nptTime))
	endOnly := regexp.MustCompile(fmt.Sprintf("(-%s)", end))
	_ = endOnly
	nptRange := regexp.MustCompile(fmt.Sprintf("^npt=((%s-(%s)?)|(-%s))$", start, end, end))
	println(nptRange.String())

	res := nptRange.FindStringSubmatch(input)
	fmt.Printf("res: %s\n", res)
	//fmt.Printf("res2: %s, index: %v\n", res[nptRange.SubexpIndex("now")], nptRange.SubexpNames())
	_ = myRange

	// check start time format // can be hh:mm:ss or X seconds
	startTimeInput := res[nptRange.SubexpIndex("start")]
	if nptSec.MatchString(startTimeInput) {
		println("range in sec")
		myRange.startTime, _ = strconv.ParseFloat(startTimeInput, 64)
	} else if nptHhmmss.MatchString(startTimeInput) {
		myRange.startTime = timeToSec(buildIsoTime(startTimeInput))
	} else {
		println("no match 1")
		myRange.liveEvent = true
	}
	// end time is optional, check if its present
	endTimeInput := res[nptRange.SubexpIndex("end")]

	if res[nptRange.SubexpIndex("end")] != "" {
		// check end time format // can be hh:mm:ss or X seconds
		if nptSec.MatchString(endTimeInput) {
			println("no match 4")
			myRange.endTime, _ = strconv.ParseFloat(endTimeInput, 64)
		} else if nptHhmmss.MatchString(endTimeInput) {
			myRange.endTime = timeToSec(buildIsoTime(endTimeInput))
		} else {
			myRange.liveEvent = true
			println("no match 2")

		}
	}
	fmt.Printf("start: %f, end: %f\n", myRange.startTime, myRange.endTime)
	fmt.Printf("range: %v\n", myRange)

	return myRange, nil
}

func (t Transport) Serialize() string {
	return ""
}

// Misc functions
func buildIsoTime(startTimeInput string) *time.Time {
	isoTimeExpr := regexp.MustCompile("[:.]")
	timeParts := isoTimeExpr.Split(startTimeInput, -1)
	var startTime time.Time

	if len(timeParts) == 3 {
		startTime = time.Date(0, 0, 0, parseInt(timeParts[0]), parseInt(timeParts[1]), parseInt(timeParts[2]), 0, time.Local)
	} else if len(timeParts) == 4 {
		float, err := strconv.ParseFloat("."+timeParts[3], 32)
		if err != nil {
			return nil
		}
		tNano := float * float64(time.Second.Nanoseconds())
		startTime = time.Date(0, 0, 0, parseInt(timeParts[0]), parseInt(timeParts[1]), parseInt(timeParts[2]), int(math.Ceil(tNano)), time.Local)
	}
	return &startTime
}

func parseInt(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
		return 0
	}
	return n

}

func timeToSec(t *time.Time) float64 {
	tSec := (float64(t.Hour()) * 60 * 60) + (float64(t.Minute()) * 60) + float64(t.Second()) + float64(t.Nanosecond())/(1000*1000*1000)
	return math.Round(tSec*1000) / 1000
}

// Removes the start of line and end of line anchors in a sub regex
func regexToString(regex *regexp.Regexp) string {
	r := strings.TrimPrefix(regex.String(), "^")
	return strings.TrimSuffix(r, "$")
}
