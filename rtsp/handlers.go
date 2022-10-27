package rtsp

import (
	"context"
	"encoding/base64"
	"github.com/gorilla/mux"
	"github.com/pion/sdp"
	"golang.org/x/exp/rand"
	"io"
	"log"
	"net/http"
	"tmss/rtsp/headers"
)

const (
	SessionLen     = 15
	rtspSessionKey = "rtsp_session"
)

func DefaultTransport() headers.Transport {
	return headers.Transport{}
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

	transports := headers.ParseTransport(request.Header.Get(TransportHeader))
	handler.sessions[sessionId] = Session{
		Transport: transports[0],
	}
	resWriter.Header().Add(TransportHeader, DefaultTransport().Serialize())
	resWriter.Header().Add(CSeqHeader, request.Header.Get(CSeqHeader))

	_, err := resWriter.Write([]byte("body"))
	if err != nil {
		log.Fatal("Failed to send res to client")
		return
	}
}

func (handler Handler) AnnounceHandler(resWriter http.ResponseWriter, request *http.Request) {
	session, found := handler.sessions[request.Header.Get("Session")]
	if !found {
		//TODO send an error msg to client
		return
	}
	desc := sdp.SessionDescription{}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}

	err = desc.Unmarshal(string(body))
	if err != nil {
		return
	}
	_ = session

	//TODO implement later
}

func (handler Handler) DescribeHandler(resWriter http.ResponseWriter, request *http.Request) {
	//TODO get video desc
	sessionDescription := &sdp.SessionDescription{}
	descriptionRaw := sessionDescription.Marshal()

	resWriter.Header().Add(CSeqHeader, request.Header.Get(CSeqHeader))
	//resWriter.Header().Add(ContentLengthHeader,request.Header.Get(CSeqHeader))
	resWriter.Header().Add(SessionHeader, request.Header.Get(SessionHeader))

	_, err := resWriter.Write([]byte(descriptionRaw))
	if err != nil {
		log.Fatalf("cannot send res,%v\n", err)
		return
	}
}

func (handler Handler) PlayHandler(resWriter http.ResponseWriter, request *http.Request) {
	var rangeHeader *headers.Range
	session := request.Context().Value("rtsp_session").(Session)

	if rangerHeaderString := request.Header.Get(RangeHeader); rangerHeaderString != "" {
		header, err := headers.ParseRange(rangerHeaderString)
		if err == nil {
			rangeHeader = &header
		}
	}
	session.Play(rangeHeader)

	// we should check if the session is still active
	_, err := resWriter.Write([]byte{})
	if err != nil {
		return
	}
}

// adds a rtsp session in http.Request;
//
//it returns a http error if the session header
// is not present and the req is not a SETUP request
func (handler Handler) setSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Proto != SETUP {
			sessionID := request.Header.Get(SessionHeader)
			if sessionID == "" {
				writer.WriteHeader(http.StatusUnauthorized)
				http.Error(writer, "missing session", http.StatusBadRequest)
				return
			}
			session := handler.sessions[sessionID]
			request.WithContext(context.WithValue(request.Context(), rtspSessionKey, session))
		}
		next.ServeHTTP(writer, request)
	})
}

/*
func (handler Handler) PauseHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) TearDownHandler(request RtspRequest, resWriter ResponseWriter) {

}
*/
