package rtsp

import (
	"errors"
	"fmt"
	"golang.org/x/exp/rand"
	"net"
	"strings"
	"time"
	"tmss/media"
	"tmss/rtsp/headers"
)

const (
	RTPProfile = "RTP/AVP"
	minPort    = 1024
	maxPort    = 65535
)

type PlayPauseRequest struct {
}

var UnsupportedProfileErr = errors.New("unsupported profile")

type Session struct {
	TransmissionType string
	AudioPort        int
	VideoPort        int
	Transport        headers.Transport
	queuePlayRequest chan headers.Range
	deQueue          chan *headers.Range
	pauseRequest     chan headers.Range
	plays            []headers.Range
	playsWatcher     chan bool
	currentRange     headers.Range
	resumePoint      float64
	streams          []MediaStreamer
}

type MediaProvider interface {
	init(mediaID string)
	get(mediaId string) string
}

type MediaStreamer interface {
	Play(timeRange headers.Range)
	Pause(timeRange headers.Range)
	HandleRtcp()
	HandleRtp()
}

func OpenNewSession(mediaId string, mediaRecord media.Media) Session {
	return Session{
		streams: make([]MediaStreamer, len(mediaRecord.Streams)),
	}
}

func (session Session) SessionStart() error {
	if !strings.Contains(session.Transport.Profile, RTPProfile) {
		return UnsupportedProfileErr
	}
	return nil
}

func (session Session) PlayPause(pause bool, timeRange headers.Range) {
	for {
		select {
		// handle pause requests
		case pause := <-session.pauseRequest:
			if pause.StartTime < session.currentRange.StartTime || pause.StartTime > session.currentRange.EndTime {
				// todo RETURN ERROR bc PAUSE time is outside of any queued PLAY range
			}
			// send the pause command to the streamer
			session.plays = nil
			//session.streams[0].getCommandChannel() <- pause
		// queue play commands and notify other waiting go routines
		case play := <-session.queuePlayRequest:
			session.plays = append(session.plays, play)
			select {
			case session.playsWatcher <- true:
			default:
			}
		// dequeue a play command and send it to streamer to send to the client
		case _ = <-session.deQueue:
			// should check if there's any play request available
			session.deQueue <- &session.plays[0]
			session.plays = session.plays[1:]
		}
	}
}

func (session Session) Play(streamRange *headers.Range) {
	return
	if streamRange == nil {
		streamRange = &headers.Range{
			StartTime: session.resumePoint,
		}
	}
	session.queuePlayRequest <- *streamRange
}

func (session Session) queueFrame() {
	for {

		session.deQueue <- &headers.Range{}
		r := <-session.deQueue

		if r == nil {
			<-session.playsWatcher
			continue
		}

		//session.streams[0].Play(*r)
	}
}

type ConnAlloc struct {
	Conn net.PacketConn
	Port int
}

func GetConn() (ConnAlloc, ConnAlloc, error) {
	startTime := time.Now()
	for time.Now().Sub(startTime) <= time.Second*30 {
		i := rand.Intn(maxPort+minPort) + minPort
		if i%2 != 0 {
			continue
		}
		rtpConn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", i))
		if err != nil {
			continue
		}
		i++
		rtcpConn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", i))
		if err != nil {
			continue
		}

		return ConnAlloc{
				Conn: rtpConn,
				Port: i - 1,
			},
			ConnAlloc{
				Conn: rtcpConn,
				Port: i,
			}, nil
	}
	return ConnAlloc{}, ConnAlloc{}, errors.New("failed to open connections")
}
