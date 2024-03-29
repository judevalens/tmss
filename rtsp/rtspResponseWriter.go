//go:generate mockgen -destination=mocks/conn.go . Conn
package rtsp

import (
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

type ResponseWriter struct {
	*http.Response
	conn        net.Conn
	isHeaderSet bool
	Headers     http.Header
}

func (r ResponseWriter) Header() http.Header {
	return r.Response.Header
}

func (r ResponseWriter) Write(data []byte) (int, error) {
	if !r.isHeaderSet {
		r.WriteHeader(http.StatusOK)
	}
	if data == nil {
		data = []byte{}
	}
	r.ContentLength = int64(len(data))
	r.Response.Body = io.NopCloser(bytes.NewReader(data))
	if r.ContentLength > 0 {
		r.Header().Add(ContentLengthHeader, strconv.Itoa(int(r.ContentLength)))

	}
	rawResponse, err := serializeResponse(*r.Response)
	if err != nil {
		log.Fatal(err)
	}
	n, err := r.conn.Write([]byte(rawResponse))
	return n, err
}

func (r ResponseWriter) WriteHeader(statusCode int) {
	r.Response.StatusCode = statusCode
	r.Response.Status = http.StatusText(statusCode)
}
