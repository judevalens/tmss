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

func InitRtpStream(media media.Media, streamId int, rtpConn net.PacketConn, rtcpConn net.PacketConn) (MediaStreamer, error) {
	var streamer MediaStreamer
	switch media.Streams[streamId].RtpFormat {
	case H264Codec:
		streamer, _ = h264.Init(media, streamId, rtpConn, rtcpConn)
	default:
		return nil, errors.New("codec is not supported")
	}
	return streamer, nil
}
