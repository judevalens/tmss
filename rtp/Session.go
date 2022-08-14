package rtp

import "net"

type Session struct {
	TransmissionType string
	AudioPort        int
	VideoPort        int
}

func OpenNewSession(mediaId string, addr net.Addr) Session {
	return Session{}
}
