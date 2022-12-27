package main

import (
	"fmt"
	"io/ioutil"
	"tmss/media"
	"tmss/rtp"
	"tmss/rtsp"
)

const mediaPath = "/home/jude/Desktop/amnis_server/big_buck_bunny.mp4"

func main() {
	//rtsp.InitRtspServer()

	a, b, err := rtsp.GetConn()
	if err != nil {
		return
	}
	jsonRepo := media.NewJsonRepo()

	mediaRecord := jsonRepo.GetMedia("44e2a14a22")
	streamer, _ := rtp.InitRtpStream(mediaRecord, 0, a.Conn, b.Conn)
	streamer.HandleRtp()
	for {

	}
	return

	file, err := ioutil.ReadFile(mediaPath)
	if err != nil {
		return
	}
	fmt.Println(len(file))
	jsonRepo = media.NewJsonRepo()
	jsonRepo.AddMedia(file, "big_buck_bunny_2.mp4")
	session := jsonRepo.GetSDPSession("269e0119ff3f875b09098ab743fe4cc7c431926febfe4bf20e02954b1c74136d")
	fmt.Printf("%v\n", session.Marshal())
}
