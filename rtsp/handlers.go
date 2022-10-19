package rtsp

import (
	"encoding/base64"
	"golang.org/x/exp/rand"
)

const SessionLen = 15

type Session struct {
}

type Handler struct {
	sessions map[string]Session
}

func (handler Handler) SetUpHandler(request RtspRequest, resWriter ResponseWriter) {
	sessionBuf := make([]byte, SessionLen)
	mediaId := request.Uri.Query().Get("media_id")

	if len(mediaId) == 0 {
		//TODO handle missing media
		return
	}

	sessionId, found := request.Headers[SessionHeader]
	if !found {
		_, err := rand.Read(sessionBuf)
		if err != nil {
			return
		}
		sessionId = base64.URLEncoding.EncodeToString(sessionBuf)
	}

	//TODO we probably should be able to create different transport based on the client

}
func (handler Handler) AnnounceHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) DescribeHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) PlayHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) PauseHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) TearDownHandler(request RtspRequest, resWriter ResponseWriter) {

}
