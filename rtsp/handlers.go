package rtsp

import (
	"encoding/base64"
	"github.com/pion/sdp"
	"golang.org/x/exp/rand"
	"log"
	"net/http"
	"strconv"
)

const SessionLen = 15

func DefaultTransport() Transport {
	return Transport{}
}

type Handler struct {
	sessions map[string]Session
}

func (handler Handler) SetUpHandler(request Request, resWriter ResponseWriter) {
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
	_ = sessionId

	transports := ParseTransport(request.Headers["Transport"])
	handler.sessions[sessionId] = Session{
		Transport: transports[0],
	}

	res := Response{
		Version:    RtspVersion,
		StatusCode: 200,
		Reason:     "OK",
	}
	res.Headers[TransportHeader] = DefaultTransport().Serialize()
	res.Headers[CSeqHeader] = request.Headers[CSeqHeader]

	_, err := resWriter.conn.Write([]byte(serializeResponse(res)))
	if err != nil {
		log.Fatal("Failed to send res to client")
		return
	}
}

func (handler Handler) AnnounceHandler(request Request, resWriter ResponseWriter) {
	session, found := handler.sessions[request.Headers["Session"]]
	if !found {
		//TODO send an error msg to client
		return
	}
	desc := sdp.SessionDescription{}
	err := desc.Unmarshal(string(request.Body))
	if err != nil {
		return
	}
	_ = session

	//TODO implement later
}

func (handler Handler) DescribeHandler(request Request, resWriter ResponseWriter) {
	//TODO get video desc
	sessionDescription := &sdp.SessionDescription{}
	descriptionRaw := sessionDescription.Marshal()

	res := Response{
		StatusCode: http.StatusOK,
		Reason:     http.StatusText(http.StatusOK),
		Version:    RtspVersion,
		Headers: map[string]string{
			CSeqHeader:          request.Headers[CSeqHeader],
			ContentLengthHeader: strconv.Itoa(len(descriptionRaw)),
		},
		Body: []byte(descriptionRaw),
	}
	_, err := resWriter.conn.Write([]byte(serializeResponse(res)))
	if err != nil {
		log.Fatalf("cannot send res,%v\n", err)
		return
	}
}

/*
func (handler Handler) PlayHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) PauseHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) TearDownHandler(request RtspRequest, resWriter ResponseWriter) {

}
*/
