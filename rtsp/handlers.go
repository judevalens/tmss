package rtsp

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pion/sdp"
	"golang.org/x/exp/rand"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
	"tmss/media"
	"tmss/rtp"
	"tmss/rtsp/headers"
)

const (
	SessionLen     = 15
	rtspSessionKey = "rtsp_session"
)

type Handler struct {
	sessions            map[string]Session
	highestPort         int
	portAssignmentMutex *sync.Mutex
	MediaRepo           media.RepoI
}

func (handler Handler) defaultTransport(sessionId string) headers.Transport {
	session := handler.sessions[sessionId]
	handler.portAssignmentMutex.Lock()
	//	handler.highestPort = session.InitServers(handler.highestPort)
	_ = session
	handler.portAssignmentMutex.Unlock()
	return headers.Transport{}
}

func (handler Handler) SetUpHandler(resWriter http.ResponseWriter, request *http.Request) {
	fmt.Printf("new set up request: %v\n", request.Method)
	sessionBuf := make([]byte, SessionLen)
	mediaId := mux.Vars(request)["id"]
	streamId, err := strconv.Atoi(mux.Vars(request)["streamId"])

	if err != nil {
		log.Fatal("invalid stream id")
		resWriter.WriteHeader(http.StatusBadRequest)
	}
	var sessionId string
	var session Session
	r := rand.New(rand.NewSource(uint64(time.Now().Nanosecond())))
	if request.Context().Value(rtspSessionKey) == nil {
		if sessionId == "" {
			_, err := r.Read(sessionBuf)
			if err != nil {
				return
			}
			sessionId = base64.URLEncoding.EncodeToString(sessionBuf)
		}
		mediaInfo := handler.MediaRepo.GetMedia(mediaId)
		session = OpenNewSession(mediaId, mediaInfo)
	} else {
		sessionId = request.Context().Value(rtspSessionKey).(string)
		session = handler.sessions[sessionId]
	}
	_ = session

	rtpConn, rtcpConn, err := GetConn()
	if err != nil {
		log.Fatal(err)
		return
	}
	mediaData := handler.MediaRepo.GetMedia(mediaId)
	//TODO check that the mediaId is valid
	session.streams[streamId], _ = rtp.InitRtpStream(mediaData, streamId, rtpConn.conn, rtcpConn.conn)
	transports := headers.ParseTransport(request.Header.Get(TransportHeader))
	handler.sessions[sessionId] = Session{
		Transport: transports[0],
	}
	//
	resWriter.Header().Add(TransportHeader, request.Header.Get(TransportHeader)+";server_port=9002-9003;ssrc=1234ABCD")
	resWriter.Header().Add(CSeqHeader, request.Header.Get(CSeqHeader))
	resWriter.Header().Add(SessionHeader, sessionId)

	_, err = resWriter.Write([]byte{})
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
func (handler Handler) OptionsHandler(resWriter http.ResponseWriter, request *http.Request) {
	fmt.Printf("new option request: %v\n", request.Method)
	resWriter.Header()
	resWriter.Header().Add(PublicHeader, fmt.Sprintf("%s, %s, %s, %s, %s", DESCRIBE, SETUP, TEARDOWN, PLAY, PAUSE))
	resWriter.Header().Add("CSeq", request.Header.Get(CSeqHeader))
	resWriter.WriteHeader(http.StatusOK)
	_, err := resWriter.Write([]byte{})
	if err != nil {
		log.Fatal(err)
		return
	}

}
func (handler Handler) DescribeHandler(resWriter http.ResponseWriter, request *http.Request) {
	//TODO get video desc
	mediaId := mux.Vars(request)["id"]
	sessionDescription := handler.MediaRepo.GetSDPSession(mediaId)
	descriptionRaw := sessionDescription.Marshal()
	resWriter.Header().Add(CSeqHeader, request.Header.Get(CSeqHeader))
	_, err := resWriter.Write([]byte(descriptionRaw))
	if err != nil {
		log.Fatalf("cannot send res,%v\n", err)
		return
	}
}
func (handler Handler) PlayHandler(resWriter http.ResponseWriter, request *http.Request) {
	resWriter.Header().Add(CSeqHeader, request.Header.Get(CSeqHeader))
	resWriter.Header().Add("Range", "npt=0.000-")

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
		if request.Method != SETUP && request.Method != OPTIONS && request.Method != DESCRIBE {
			sessionID := request.Header.Get(SessionHeader)
			if sessionID == "" {
				writer.WriteHeader(http.StatusUnauthorized)
				http.Error(writer, "missing session", http.StatusBadRequest)
				return
			}
			session := handler.sessions[sessionID]
			request.WithContext(context.WithValue(request.Context(), rtspSessionKey, session))
		}
		println("next")
		println(next)
		next.ServeHTTP(writer, request)
	})
}

/*
func (handler Handler) PauseHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) TearDownHandler(request RtspRequest, resWriter ResponseWriter) {

}
*/
