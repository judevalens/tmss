package headers

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"tmss/rtsp/misc"
)

type Transport struct {
	Profile        string `field_type:"val"`
	LowerTransport string
	Append      string   `field_type:"key" optional:"true" name:"append"`
	CastMode    string `field_type:"val" optional:"false"`
	Destination string `field_type:"key_val" optional:"true"`
	Interleaved    string `field_type:"key_val" optional:"false"`
	TTL            string `field_type:"key_val" optional:"false"`
	Layers         string `field_type:"key_val" optional:"false"`
	Port           string `field_type:"key_val" optional:"false"`
	ClientPort     string `field_type:"key_val" optional:"true" name:"client_port"`
	ServerPort     string `field_type:"key_val" optional:"true" name:"server_port"`
	Ssrc           string `field_type:"key_val" optional:"false"`
	Mode           string `field_type:"key_val" optional:"false"`
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
		transports = append(transports, Transport{})
		components := strings.Split(transport, ";")
		transports[i].Profile = components[0]

		for j := 1; j < len(components); j++ {
			paramComponents := strings.Split(components[j], "=")
			key := strings.TrimSpace(paramComponents[0])
			hasValue := len(paramComponents) >= 2
			switch key {
			case "unicast","multicast":
				transports[i].CastMode = key
			case "append":
				transports[i].Append = paramComponents[1]
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

func SerializeTransport() {}

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
	var transportType = reflect.TypeOf(t)
	TransportR := reflect.ValueOf(t)
	serializedTransport := ""
	for i := 0; i < transportType.NumField(); i++ {
		field := transportType.Field(i)
		val :=TransportR.Field(i).String()
		key  :=  strings.ToLower(field.Name)

		if tag,found := field.Tag.Lookup("name");found {
			key	= tag
		}
		if val == "" {
			continue
		}
		switch field.Tag.Get("field_type") {
		case "val":
			serializedTransport += val
		case "key" :
			serializedTransport += key
		case "key_val":
				serializedTransport += key+"="+val
		}
		if i < transportType.NumField()-1 {
			serializedTransport +=";"
		}
	}
	fmt.Printf("%v\n",serializedTransport)
	return serializedTransport
}
