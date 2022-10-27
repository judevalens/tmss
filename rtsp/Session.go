package rtsp

import (
	"errors"
	"net"
	"strings"
)

const (
	RTPProfile = "RTP/AVP"
)

type PlayPauseRequest struct {
}

var UnsupportedProfileErr = errors.New("unsupported profile")

type Session struct {
	TransmissionType string
	AudioPort        int
	VideoPort        int
	Transport        Transport
	queuePlayRequest chan Range
	deQueue          chan *Range
	pauseRequest     chan Range
	plays            []Range
	playsWatcher     chan bool
	currentRange     Range
	resumePoint      float64
	MediaStreamer
}

type MediaStreamer interface {
	initialize(mediaID string)
	play(timeRange Range)
	pause(timeRange Range)
	getCommandChannel() chan Range
}

func OpenNewSession(mediaId string, addr net.Addr) Session {
	return Session{}
}

func (session Session) SessionStart() error {
	if !strings.Contains(session.Transport.Profile, RTPProfile) {
		return UnsupportedProfileErr
	}
	return nil
}

func (session Session) PlayPause(pause bool, timeRange Range) {
	for {
		select {
		// handle pause requests
		case pause := <-session.pauseRequest:
			if pause.startTime < session.currentRange.startTime || pause.startTime > session.currentRange.endTime {
				// todo RETURN ERROR bc PAUSE time is outside of any queued PLAY range
			}
			// send the pause command to the streamer
			session.plays = nil
			session.getCommandChannel() <- pause
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

func (session Session) Play(streamRange *Range) {
	if streamRange == nil {
		streamRange = &Range{
			startTime: session.resumePoint,
		}
	}
	session.queuePlayRequest <- *streamRange
}

func (session Session) queueFrame() {
	for {

		session.deQueue <- &Range{}
		r := <-session.deQueue

		if r == nil {
			<-session.playsWatcher
			continue
		}

		session.play(*r)
	}
}
