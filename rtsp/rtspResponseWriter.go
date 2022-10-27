package rtsp

import (
	"net"
	"net/http"
)

type ResponseWriter struct {
	*http.Response
	conn        net.Conn
	isHeaderSet bool
}

func (r ResponseWriter) Header() http.Header {
	return r.Response.Header
}

func (r ResponseWriter) Write(bytes []byte) (int, error) {
	if !r.isHeaderSet {
		r.WriteHeader(http.StatusOK)
	}
	return r.Write([]byte(serializeResponse(*r.Response)))
}

func (r ResponseWriter) WriteHeader(statusCode int) {
	r.Response.StatusCode = statusCode
	r.Response.Status = http.StatusText(statusCode)
}
