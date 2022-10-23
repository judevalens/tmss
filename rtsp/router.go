package rtsp

import (
	"fmt"
	"net"
)

const mtu = 4096

const (
	ANNOUNCE     = "ANNOUNCE"
	GetParameter = "GET_PARAMETER"
	OPTIONS      = "OPTIONS"
	PAUSE        = "PAUSE"
	PLAY         = "PLAY"
	RECORD       = "RECORD"
	REDIRECT     = "REDIRECT"
	SETUP        = "SETUP"
	SetParameter = "SET_PARAMETER"
	TEARDOWN     = "TEARDOWN"
)

type Router struct {
	name    string
	Paths   map[string]*Router
	handler RequestHandler
}
type ResponseWriter struct {
	conn net.Conn
}
type RequestHandler func(request Request, rtspWriter ResponseWriter)

func NewRouter() Router {
	return Router{
		Paths: map[string]*Router{},
	}
}
func Print(paths map[string]*Router, level int) {
	offset := ""

	for i := 0; i < level; i++ {
		offset += " "
	}
	for pathP, pathS := range paths {
		fmt.Printf("%v|-> %v\n", offset, pathP)
		Print(pathS.Paths, len(pathP)+level+3)
	}
}

/*
func (router RtspRouter) AddRoute(method string, path string, handler RequestHandler) {
	path = method + path
	paths := strings.Split(path, "/")
	router.add(paths, handler)
}
func (router RtspRouter) add(path []string, handler RequestHandler) {
	rtspPath, found := router.Paths[path[0]]
	if !found {
		rtspPath = &RtspRouter{
			Paths: map[string]*RtspRouter{},
		}
		router.Paths[path[0]] = rtspPath
		fmt.Printf("adding path: %v\n", path[0])
	}
	rtspPath.name = path[0]
	if len(path) == 1 {
		rtspPath.handler = handler
		return
	}
	rtspPath.add(path[1:], handler)
}
func (router RtspRouter) route(request RtspRequest, resWriter ResponseWriter) {
	paths := strings.Split(request.Method+request.Uri.Path, "/")
	handler := router.findHandler(paths)
	if handler != nil {
		handler(request, resWriter)
	}
}
func (router RtspRouter) findHandler(paths []string) RequestHandler {
	if len(paths) == 1 && router.name == paths[0] && router.handler != nil {
		return router.handler
	}
	r, found := router.Paths[paths[0]]
	if found {
		return r.findHandler(paths[1:])
	}
	return nil
}
func StartServer(addr net.Addr, router RtspRouter) {
	listener, err := net.Listen(addr.Network(), addr.String())
	if err != nil {
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		//TODO might be wise to use queue
		go handleNewReq(conn, router)
	}
}

func handleNewReq(conn net.Conn, router RtspRouter) {
	buff := make([]byte, mtu)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Printf("new rtsp message, size: %v\n", n)
		request, err := ParseRequest(string(buff[:n]))
		if err != nil {
			log.Printf("Failed to parse rtsp request")
			continue
		}
		if request.BodySize > 0 {
			request.Body = string(buff[request.headerSize : request.headerSize+request.BodySize])
		}
		//check if There are still more data over the wire .... meaning data sent is larger than Buffer
		if request.BodySize+request.headerSize > n {
			//TODO we should handle that
		}
		resWriter := ResponseWriter{
			conn: conn,
		}
		router.route(request, resWriter)
	}
}
func HandleNewReq2(request RtspRequest, rtspWriter ResponseWriter) {
}
*/
