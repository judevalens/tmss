package media

import (
	"github.com/pion/sdp"
	"time"
)

const utcDiff = (70*365 + 17) * 24 * 60 * 60

type RepoI interface {
	GetSDPSession(mediaId string) *sdp.SessionDescription
}

type JsonRepo struct {
}

func (repo JsonRepo) GetSDPSession(mediaId string) *sdp.SessionDescription {
	session := &sdp.SessionDescription{
		Version: sdp.Version(0),
		//TODO find a solution to get external ip addr
		Origin: sdp.Origin{
			Username:       "-",
			SessionID:      utcToNtp(time.Now().UnixMilli()),
			SessionVersion: utcToNtp(time.Now().UnixMilli()),
			NetworkType:    "IN",
			AddressType:    "IPv4",
			UnicastAddress: "localhost",
		},
	}
	session.SessionName = "video streaming"
	session.MediaDescriptions[0] = &sdp.MediaDescription{}
	return session
}

func ntpToUtc(ts uint64) int64 {
	seconds := (ts >> 32) & 0xFFFFFFFF
	milliSeconds := ts & 0xFFFFFFFF
	u := seconds - utcDiff
	u += milliSeconds / 1000
	return int64(u)
}

func utcToNtp(ts int64) uint64 {
	var seconds uint64
	seconds = seconds & uint64(ts)
	seconds <<= 32
	return seconds
}
