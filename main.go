package main

import (
	"fmt"
	"io/ioutil"
	"tmss/media"
)

const mediaPath = "/home/jude/Desktop/amnis_server/big_buck_bunny.mp4"

func main() {
	//rtsp.InitRtspServer()

	file, err := ioutil.ReadFile(mediaPath)
	if err != nil {
		return
	}
	fmt.Println(len(file))
	jsonRepo := media.NewJsonRepo()
	jsonRepo.AddMedia(file, "big_buck_bunny_2.mp4")
	session := jsonRepo.GetSDPSession("269e0119ff3f875b09098ab743fe4cc7c431926febfe4bf20e02954b1c74136d")
	fmt.Printf("%v\n", session.Marshal())
}
