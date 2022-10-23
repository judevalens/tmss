package rtsp

import (
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var body string
var invalidReq string
var req string
var reader io.Reader

func setUp() {
	body = "322205525053950747749964020551555970594761125490594877568263165237635119602792619464638218369285983252070845171520636711203132229485425590254261149928915273764367882354151353178108828292366131332677437982903377233132407844524463374388033395654714610115003331494340674005390301680361719064451717414907362014946063028211777331341212210517278080888638017317238906999197122070512658445773911885201499350537989656582088496876487307411531773758716679718662138334615863268258290679213915327280954751009151387033720252059413883252967122160809677943274943879481085154486194005629450569978894259442807466344993005061929056547114513563459522342755211412107637811865370245304799508604319298883809851913759934756776378246555505546219561823439410977440326920676022028468173315359739249472036964130027054788847620436164598258356890082971941650075231523703703883969235623680273106402461937007904809308963535793182094301134962389981680279101121925641158294216521656399161807448495037792049975124075384760686521403531863283249499732842263260513669117580337156604606105692129261461828329516885315787849935490628497448861940994983677732575492053427805093746557362291733542608703716239239080002126074310259655938677310904832529095186063247868960800524873"

	//DO NOT CHANGE, MUST IDENTICAL TO PASS CERTAIN TESTS
	req = "" +
		"SETUP rtsp://example.com/media.mp4/streamid=0 RTSP/1.0\r\n" +
		"CSeq: 3\r\n" +
		"Content-Length: " + strconv.Itoa(len(body)) + "\r\n" +
		"Transport: RTP/AVP;multicast;ttl=127;mode=\"PLAY\",RTP/AVP;unicast;client_port=3456-3457;mode=\"PLAY\"\r\n" +
		"\r\n" +
		body

	invalidReq = "SETUP rtsp://example.com/media.mp4/streamid=0 RTSP/1.0" +
		"CSeq: 3\r\n" +
		"Content-Length: " + strconv.Itoa(len(body)) + "\r\n" +
		"Transport: RTP/AVP;multicast;ttl=127;mode=\"PLAY\",RTP/AVP;unicast;client_port=3456-3457;mode=\"PLAY\"\r\n" +
		"\r\n"
}

func TestParseRequest(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Request
		wantErr bool
	}{
		{
			name: "Test parsing valid rtsp request",
			want: Request{
				Method:  "SETUP",
				Version: "RTSP/1.0",
				Uri: func() *url.URL {
					parse, err := url.Parse("rtsp://example.com/media.mp4/streamid=0")
					if err != nil {
						return nil
					}
					return parse
				}(),
				Headers: func() map[string]string {
					setUp()
					return map[string]string{
						"CSeq":           "3",
						"Transport":      "RTP/AVP;multicast;ttl=127;mode=\"PLAY\",RTP/AVP;unicast;client_port=3456-3457;mode=\"PLAY\"",
						"Content-Length": strconv.Itoa(len(body)),
					}
				}(),
				Body: []byte(body),
			},
			args: args{
				reader: func() io.Reader {
					setUp()
					return strings.NewReader(req)
				}(),
			},
		},
		{
			name: "Test parsing rtsp request with invalid statusLine",
			want: Request{
				Method:  "SETUP",
				Version: "RTSP/1.0",
				Uri: func() *url.URL {
					parse, err := url.Parse("rtsp://example.com/media.mp4/streamid=0")
					if err != nil {
						return nil
					}
					return parse
				}(),
				Headers: func() map[string]string {
					setUp()
					return map[string]string{
						"CSeq":           "3",
						"Transport":      "RTP/AVP;multicast;ttl=127;mode=\"PLAY\",RTP/AVP;unicast;client_port=3456-3457;mode=\"PLAY\"",
						"Content-Length": strconv.Itoa(len(body)),
					}
				}(),
				Body: []byte(body),
			},
			args: args{
				reader: func() io.Reader {
					setUp()
					return strings.NewReader(invalidReq)
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRequest(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
