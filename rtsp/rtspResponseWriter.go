//go:generate mockgen -destination=mocks/conn.go . Conn
package rtsp

import (
	"bytes"
	"io"
	"net"
	"net/http"
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
}

func (r ResponseWriter) Header() http.Header {
	return r.Response.Header
}

func (r ResponseWriter) Write(data []byte) (int, error) {
	if !r.isHeaderSet {
		r.WriteHeader(http.StatusOK)
	}
	r.Response.Body = io.NopCloser(bytes.NewReader(data))
	return r.conn.Write([]byte(serializeResponse(*r.Response)))
}

func (r ResponseWriter) WriteHeader(statusCode int) {
	r.Response.StatusCode = statusCode
	r.Response.Status = http.StatusText(statusCode)
}
