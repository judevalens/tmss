package rtp

import (
	"errors"
	"net"
	"tmss/media"
	"tmss/rtp/h264"
	"tmss/rtsp/headers"
)

type MediaControl struct {
	teardown bool
}

const (
	H264Codec = "H264"
	AccCodec  = "MP4A-LATM"
)

type MediaStreamer interface {
	Play(timeRange headers.Range)
	Pause(timeRange headers.Range)
	HandleRtcp()
	HandleRtp()
}

type Session struct {
	AudioPort    int
	VideoPort    int
	mediaId      string
	mediaControl chan MediaControl
}

func (session Session) HandleRtcp() {

}
func (session Session) HandleRtp() {

}
func (session Session) Play(timeRange headers.Range) {

}

func (session Session) Pause(timeRange headers.Range) {

}

func InitRtpStream(media media.Media, streamId int, rtpConn net.PacketConn, rtcpConn net.PacketConn) (MediaStreamer, error) {
	var streamer MediaStreamer
	switch media.Streams[streamId].RtpFormat {
	case H264Codec:
		streamer, _ = h264.Init(media, streamId, rtpConn, rtcpConn)
	case AccCodec:
		return Session{}, nil
	default:
		return nil, errors.New("codec is not supported")
	}
	return streamer, nil
}
