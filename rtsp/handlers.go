package rtsp

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pion/sdp"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
	"io"
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

func (handler Handler) SetUpHandler(res http.ResponseWriter, request *http.Request) {
	sessionBuf := make([]byte, SessionLen)
	mediaId := mux.Vars(request)["id"]
	streamId, err := strconv.Atoi(mux.Vars(request)["streamId"])
	if err != nil {
		zap.L().Sugar().Errorw("Invalid or missing stream id", "stream id", mux.Vars(request)["streamId"])
		res.WriteHeader(http.StatusBadRequest)
		_, err := res.Write([]byte{})
		if err != nil {
			zap.Error(err)
			return
		}
		return
	}
	zap.L().Sugar().Infow("new incoming request", "Method", request.Method, "stream id", streamId)
	var sessionId string
	var session Session
	r := rand.New(rand.NewSource(uint64(time.Now().Nanosecond())))
	if request.Context().Value(rtspSessionKey) == nil {
		if sessionId == "" {
			_, err := r.Read(sessionBuf)
			if err != nil {
				zap.L().Sugar().Error(err)
				res.WriteHeader(http.StatusInternalServerError)
				_, err := res.Write(nil)
				if err != nil {
					return
				}
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
		res.WriteHeader(http.StatusInternalServerError)
		_, err := res.Write(nil)
		if err != nil {
			return 
		}
		return
	}
	mediaData := handler.MediaRepo.GetMedia(mediaId)
	//TODO check that the mediaId is valid
	session.streams[streamId], err = rtp.InitRtpStream(mediaData, streamId, rtpConn.Conn, rtcpConn.Conn)

	if err != nil {
		zap.L().Sugar().Errorw("failed to set up stream","stream id",streamId,"err",err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(nil)
		return
	}
	session.streams[streamId].HandleRtcp()
	session.streams[streamId].HandleRtp()
	transports := headers.ParseTransport(request.Header.Get(TransportHeader))
	handler.sessions[sessionId] = Session{
		Transport: transports[0],
	}
	transportHeader := headers.ParseTransport(request.Header.Get(TransportHeader))
	transportHeader[0].ServerPort = fmt.Sprintf("%v-%v", rtpConn.Port, rtcpConn.Port)
	res.Header().Add(TransportHeader, transportHeader[0].Serialize())
	res.Header().Add(CSeqHeader, request.Header.Get(CSeqHeader))
	res.Header().Add(SessionHeader, sessionId)
	_, err = res.Write([]byte{})
	if err != nil {
		zap.L().Sugar().Errorw("failed to send res", "stream id",streamId,"remote addr",request.RemoteAddr,"err",err)
		return
	}
}

func (handler Handler) AnnounceHandler(resWriter http.ResponseWriter, request *http.Request) {
	zap.L().Sugar().Infow("new incoming request", "Method", request.Method, "remote addr", request.RemoteAddr)
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
	zap.L().Sugar().Infow("new incoming request", "Method", request.Method, "remote addr", request.RemoteAddr)
	resWriter.Header()
	resWriter.Header().Add(PublicHeader, fmt.Sprintf("%s, %s, %s, %s, %s", DESCRIBE, SETUP, TEARDOWN, PLAY, PAUSE))
	resWriter.Header().Add("CSeq", request.Header.Get(CSeqHeader))
	resWriter.WriteHeader(http.StatusOK)
	_, err := resWriter.Write([]byte{})
	if err != nil {
		zap.L().Sugar().Error(err)
		return
	}
}
func (handler Handler) DescribeHandler(resWriter http.ResponseWriter, request *http.Request) {
	zap.L().Sugar().Infow("new incoming request", "Method", request.Method, "remote addr", request.RemoteAddr)
	mediaId := mux.Vars(request)["id"]
	sessionDescription := handler.MediaRepo.GetSDPSession(mediaId)
	descriptionRaw := sessionDescription.Marshal()
	resWriter.Header().Add(CSeqHeader, request.Header.Get(CSeqHeader))
	_, err := resWriter.Write([]byte(descriptionRaw))
	if err != nil {
		zap.L().Sugar().Error(err)
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
		next.ServeHTTP(writer, request)
	})
}

/*
func (handler Handler) PauseHandler(request RtspRequest, resWriter ResponseWriter) {

}
func (handler Handler) TearDownHandler(request RtspRequest, resWriter ResponseWriter) {

}
*/
