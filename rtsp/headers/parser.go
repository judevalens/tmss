package headers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"tmss/rtsp/misc"
)

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
	StartTime float64
	EndTime   float64
	Live      bool
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
	var nptTime = regexp.MustCompile(fmt.Sprintf("((?P<now>now)|%s|%s)", misc.RegexToString(nptHhmmss), misc.RegexToString(nptSec)))
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
		myRange.StartTime, _ = strconv.ParseFloat(startTimeInput, 64)
	} else if nptHhmmss.MatchString(startTimeInput) {
		myRange.StartTime = misc.TimeToSec(misc.BuildIsoTime(startTimeInput))
	} else {
		println("no match 1")
		myRange.Live = true
	}
	// end time is optional, check if its present
	endTimeInput := res[nptRange.SubexpIndex("end")]

	if res[nptRange.SubexpIndex("end")] != "" {
		// check end time format // can be hh:mm:ss or X seconds
		if nptSec.MatchString(endTimeInput) {
			println("no match 4")
			myRange.EndTime, _ = strconv.ParseFloat(endTimeInput, 64)
		} else if nptHhmmss.MatchString(endTimeInput) {
			myRange.EndTime = misc.TimeToSec(misc.BuildIsoTime(endTimeInput))
		} else {
			myRange.Live = true
			println("no match 2")

		}
	}
	fmt.Printf("start: %f, end: %f\n", myRange.StartTime, myRange.EndTime)
	fmt.Printf("range: %v\n", myRange)

	return myRange, nil
}

func (t Transport) Serialize() string {
	return ""
}
