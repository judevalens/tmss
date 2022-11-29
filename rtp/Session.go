package rtp

import (
	"net"
	"tmss/rtsp/headers"
)

const MTU = 65000

type MediaControl struct {
	teardown bool
}

type MediaStreamer interface {
	Play(timeRange headers.Range)
	Pause(timeRange headers.Range)
	Init(mediaId string, rtpConn net.PacketConn, rtcpConn net.PacketConn)
}

type Session struct {
	AudioPort    int
	VideoPort    int
	mediaId      string
	mediaControl chan MediaControl
}

func (session Session) Play(timeRange headers.Range) {
	//TODO implement me
	panic("implement me")
}

func (session Session) Pause(timeRange headers.Range) {
	//TODO implement me
	panic("implement me")
}

func (session Session) Init(mediaId string, rtpConn net.PacketConn, rtcpConn net.PacketConn) {
	//TODO implement me
	panic("implement me")
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
