package media

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "lib/mms/video_reader.h"
import "C"
import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pion/sdp"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
	misc2 "tmss/misc"
)

const Transport = "RTP/AVP"

type RepoI interface {
	GetSDPSession(mediaId string) *sdp.SessionDescription
}

type JsonRepo struct {
	media map[string]Media
}

type Media struct {
	ID        string
	Name      string `json:"name:name"`
	MediaType string
	Duration  int64
	Format    string
	Streams   []Stream
}
type Stream struct {
	MediaType string
	Format    string
}

func (repo JsonRepo) GetSDPSession(mediaId string) *sdp.SessionDescription {
	media := repo.media[mediaId]
	session := &sdp.SessionDescription{
		Version: sdp.Version(0),
		//TODO find a solution to get external ip addr
		Origin: sdp.Origin{
			Username:       "-",
			SessionID:      misc2.UtcToNtp(time.Now().UnixMilli()),
			SessionVersion: misc2.UtcToNtp(time.Now().UnixMilli()),
			NetworkType:    "IN",
			AddressType:    "IPv4",
			UnicastAddress: "localhost",
		},
	}
	session.SessionName = "video streaming"
	session.MediaDescriptions[0] = &sdp.MediaDescription{}
	session.MediaDescriptions = make([]*sdp.MediaDescription, len(media.Streams))
	for i, stream := range media.Streams {
		session.MediaDescriptions[i] = &sdp.MediaDescription{
			MediaName: sdp.MediaName{
				Media:   stream.MediaType,
				Protos:  []string{Transport},
				Formats: []string{getRTPPayloadType(stream.Format)},
			},
		}
	}
	return session
}
func (repo JsonRepo) AddMedia(rawMedia []byte, name string) {
	hash := sha256.New()
	mediaID := hex.EncodeToString(hash.Sum(rawMedia))
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	mediaPath := path.Join(homeDir, "Desktop/amnis_server")
	err = os.MkdirAll(mediaPath, os.ModeDir)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ioutil.WriteFile(path.Join(mediaPath, name+mediaID), rawMedia, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func getRTPPayloadType(format string) string {
	return ""
}
