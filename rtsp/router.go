package rtsp

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"path"
	"tmss/media"
)

const mtu = 4096

func InitRtspServer() {
	handler := Handler{
		sessions:  map[string]Session{},
		MediaRepo: media.NewJsonRepo(),
	}
	StartServer(newRouter(handler))
}

func newRouter(handler Handler) *mux.Router {
	defaultPath := "/media/{id}"
	r := mux.NewRouter()
	r.HandleFunc(path.Join(defaultPath, "{streamId}"), handler.SetUpHandler).Methods(SETUP)
	r.HandleFunc(defaultPath, handler.OptionsHandler).Methods(OPTIONS)
	r.HandleFunc(defaultPath, handler.DescribeHandler).Methods(DESCRIBE)
	r.HandleFunc(defaultPath, handler.PlayHandler).Methods(PLAY)
	r.Use(handler.setSession)
	return r
}

func StartServer(router *mux.Router) {
	tcp, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 5000})
	if err != nil {
		log.Fatal(err)
		return
	}
	zap.L().Sugar().Infow("RTSP server is running","addr",tcp.Addr())
	for {

		acceptTCP, err := tcp.AcceptTCP()
		if err != nil {
			zap.L().Sugar().Error(err)
			return
		}
		// accept and handle a new requests

		go func(conn *net.TCPConn) {
			for {
				request, err := ParseRequest(conn)
				if err != nil {
					zap.L().Sugar().Errorw("failed to parse request","client addr",conn.RemoteAddr())
					//log.Fatalf("failed to parse request\n%v", err)
					return
				}
				router.ServeHTTP(ResponseWriter{
					Response: &http.Response{
						Proto:  RtspVersion,
						Header: map[string][]string{},
					},
					conn: acceptTCP,
				}, request)
			}

		}(acceptTCP)
	}

}
