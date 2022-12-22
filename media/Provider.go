package media

import "C"
import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pion/sdp"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media -lavcodec -lavutil
// #include "/usr/local/usr/include/video_reader.h"
import "C"

const Transport = "RTP/AVP"
const JsonRepoName = "media.json"
const AppDir = "Desktop/amnis_server"

type Context C.struct_AVFormatContext

type RepoI interface {
	GetSDPSession(mediaId string) *sdp.SessionDescription
	GetMedia(mediaId string) Media
}

type JsonRepo struct {
	Media map[string]Media
}

type Media struct {
	ID        string
	Name      string
	InputName string
	MimeType  string
	Format    string
	Streams   []Stream
}
type Stream struct {
	Path        string
	MediaType   string
	Format      string
	RtpFormat   string
	ClockRate   int
	Duration    int64
	TimeBaseNum int
	TimeBaseDen int
	PayloadType int
	BitRate     int
	SampleRate  int
	FPS         int
}

func NewJsonRepo() *JsonRepo {
	repo := &JsonRepo{
		map[string]Media{},
	}
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}
	repoPath := path.Join(homeDir, AppDir, JsonRepoName)
	file, err := ioutil.ReadFile(repoPath)
	if err != nil {
		//log.Fatal(err)
		return repo
	}
	err = json.Unmarshal(file, repo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return repo
}

func (repo JsonRepo) GetSDPSession(mediaId string) *sdp.SessionDescription {
	media := repo.Media[mediaId]
	session := &sdp.SessionDescription{
		Version: sdp.Version(0),
		//TODO find a solution to get external ip addr
		Origin: sdp.Origin{
			Username:       "-",
			SessionID:      uint64(time.Now().UnixMilli()),
			SessionVersion: uint64(time.Now().UnixMilli()),
			NetworkType:    "IN",
			AddressType:    "IPv4",
			UnicastAddress: "localhost",
		},
	}
	session.SessionName = "video streaming"
	session.MediaDescriptions = make([]*sdp.MediaDescription, len(media.Streams))
	port := 5090
	for i, stream := range media.Streams {
		session.MediaDescriptions[i] = &sdp.MediaDescription{
			MediaName: sdp.MediaName{
				Media:   stream.MediaType,
				Port:    sdp.RangedPort{Value: 0},
				Protos:  []string{Transport},
				Formats: []string{strconv.Itoa(stream.PayloadType)},
			},
			Attributes: []sdp.Attribute{
				{
					Key:   "rtpmap",
					Value: fmt.Sprintf("%d %s/%d", stream.PayloadType, stream.RtpFormat, stream.ClockRate),
				},
				{
					Key:   "control",
					Value: fmt.Sprintf("streamid=%v", i),
				},
			},
		}
		port += 2
	}
	return session
}
func (repo JsonRepo) GetMedia(mediaId string) Media {
	return repo.Media[mediaId]
}
func (repo JsonRepo) AddMedia(rawMedia []byte, name string) {
	hash := sha1.New()
	hash.Write(rawMedia)
	var hashedId []byte
	mediaID := hex.EncodeToString(hash.Sum(hashedId)[:5])
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
	ext := path.Ext(name)
	fmt.Printf("file path: %v\n", mediaID)
	fileParts := strings.Split(path.Base(name), ext)
	fileName := fileParts[0] + mediaID + ext
	filePath := path.Join(mediaPath, fileName)
	fmt.Printf("file path: %v\n", filePath)
	err = ioutil.WriteFile(filePath, rawMedia, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	mediaCtx := C.open_media(C.CString(filePath))
	media := buildMedia(mediaCtx)
	media.Name = fileName
	repo.Media[mediaID] = media
	fmt.Printf("%v\n", media)

	streamNames := C.demux_file(mediaCtx)
	for i := 0; i < len(media.Streams); i++ {
		streamNamesPtr := unsafe.Pointer(streamNames)
		streamName := (**C.char)(unsafe.Add(streamNamesPtr, unsafe.Sizeof(streamNamesPtr)*uintptr(i)))
		media.Streams[i].Path = C.GoString(*streamName)
	}
	mediaRepoJson, err := json.Marshal(repo)
	if err != nil {
		return
	}
	mediaRepoPath := path.Join(homeDir, AppDir, JsonRepoName)
	fmt.Printf("%v", string(mediaRepoJson))
	err = ioutil.WriteFile(mediaRepoPath, mediaRepoJson, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func buildMedia(mediaCtx *C.struct_AVFormatContext) Media {
	media := Media{
		InputName: C.GoString(mediaCtx.iformat.long_name),
		MimeType:  C.GoString(mediaCtx.iformat.mime_type),
		Streams:   make([]Stream, mediaCtx.nb_streams),
	}

	var i C.uint = 0
	ptr := unsafe.Pointer(mediaCtx.streams)
	payloadType := 96
	for i = 0; i < mediaCtx.nb_streams; i++ {
		AVStreams := (**C.struct_AVStream)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(*mediaCtx.streams)))
		avStream := *AVStreams
		avCodec := C.avcodec_find_encoder(avStream.codecpar.codec_id)
		profile := C.av_get_profile_name(avCodec, avStream.codecpar.profile)

		fmt.Printf("codec id: %v, codec profile: %v, sample: %v\n", C.GoString(avCodec.long_name), avStream.codecpar.profile, C.GoString(profile))
		stream := Stream{
			MediaType:   C.GoString(C.av_get_media_type_string(avStream.codecpar.codec_type)),
			Format:      C.GoString(C.avcodec_get_name(avStream.codecpar.codec_id)),
			RtpFormat:   C.GoString(C.get_rtp_payload_format(avStream.codecpar.codec_id)),
			ClockRate:   int(C.get_rtp_clock_rate(avStream.codecpar.codec_id)),
			Duration:    int64(avStream.duration),
			TimeBaseNum: int(avStream.time_base.num),
			TimeBaseDen: int(avStream.time_base.den),
			PayloadType: payloadType,
			FPS:         int(avStream.avg_frame_rate.num),
			SampleRate:  int(avStream.codecpar.sample_rate),
			BitRate:     int(avStream.codecpar.bit_rate),
		}
		payloadType++
		media.Streams[i] = stream
	}

	return media
}

func GetStreamPath(streamName string) string{
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return path.Join(homeDir, AppDir,streamName)
}
