package rtp

import "net"

const MTU = 65000

type MediaControl struct {
	teardown bool
}

type Session struct {
	AudioPort    int
	VideoPort    int
	mediaId      string
	mediaControl chan MediaControl
}

func (session Session) init() {
}

func (session Session) startAudioTransmitter() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{Port: session.AudioPort})
	if err != nil {
		return
	}
	buff := make([]byte, MTU)
	_, _, err = udp.ReadFrom(buff)
	if err != nil {
		return
	}

	for {
		select {
		case mediaControl := <-session.mediaControl:
			{
				if mediaControl.teardown {
					return
				}
				//TODO transmit audio
				session.transmitAudioPacket(udp, mediaControl)
			}
		}
	}
}

func (session Session) transmitAudioPacket(conn *net.UDPConn, mediaControl MediaControl) {
	//TODO transmit data
}
