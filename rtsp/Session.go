package rtsp

import (
	"errors"
	"net"
	"strings"
)

const (
	RTPProfile = "RTP/AVP"
)

var UnsupportedProfileErr = errors.New("unsupported profile")

type Session struct {
	TransmissionType string
	AudioPort        int
	VideoPort        int
	Transport        Transport
	MediaStreamer
}
type MediaStreamer interface {
	initialize()
}

func OpenNewSession(mediaId string, addr net.Addr) Session {
	return Session{}
}

func (session Session) SessionStart() error {
	if !strings.Contains(session.Transport.Profile, RTPProfile) {
		return UnsupportedProfileErr
	}

}
