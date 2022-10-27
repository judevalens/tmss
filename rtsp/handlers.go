package rtsp

import (
	"encoding/base64"
	"github.com/gorilla/mux"
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

func (handler Handler) SetUpHandler(resWriter http.ResponseWriter, request *http.Request) {
	sessionBuf := make([]byte, SessionLen)
	mediaId := mux.Vars(request)["id"]

	if len(mediaId) == 0 {
		//TODO handle missing media
		return
	}

	sessionId := request.Header.Get(SessionHeader)
	if sessionId == "" {
		_, err := rand.Read(sessionBuf)
		if err != nil {
			return
		}
		sessionId = base64.URLEncoding.EncodeToString(sessionBuf)
	}
	_ = sessionId

	transports := ParseTransport(request.Header.Get(TransportHeader))
	handler.sessions[sessionId] = Session{
		Transport: transports[0],
	}

	res := Response{
		Version:    RtspVersion,
		StatusCode: 200,
		Reason:     "OK",
	}
	res.Headers[TransportHeader] = DefaultTransport().Serialize()
	res.Headers[CSeqHeader] = request.Header.Get(CSeqHeader)

	_, err := resWriter.Write([]byte("body"))
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

	resWriter.Response.Header = map[string][]string{
		CSeqHeader:          {request.Headers[CSeqHeader]},
		ContentLengthHeader: {strconv.Itoa(len(descriptionRaw))},
		SessionHeader:       {request.Headers[SessionHeader]},
	}
	_, err := resWriter.conn.Write([]byte(descriptionRaw))
	if err != nil {
		log.Fatalf("cannot send res,%v\n", err)
		return
	}
}

func (handler Handler) PlayHandler(request Request, resWriter ResponseWriter) {
	var rangeHeader *Range
	if rangerHeaderString, found := request.Headers[RangeHeader]; found {
		header, err := ParseRange(rangerHeaderString)
		if err == nil {
			rangeHeader = &header
		}
	}
	request.session.Play(rangeHeader)

	// we should check if the session is still active
	_, err := resWriter.Write([]byte{})
	if err != nil {
		return
	}
}

/*
func (handler Handler) PauseHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) TearDownHandler(request RtspRequest, resWriter ResponseWriter) {

}
*/
