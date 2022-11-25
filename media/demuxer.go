package media

// #cgo LDFLAGS: -L${SRCDIR}/lib/mms/build -lmms_media
// #include "../lib/mms/video_reader.h"
import "C"

type Demuxer struct {
}

func (demuxer Demuxer) process(filePath string) {

}
