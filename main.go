package main

import (
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"os"
	"tmss/media"
	"tmss/rtsp"
)

const mediaPath = "/home/jude/Desktop/amnis_server/big_buck_bunny.mp4"
const outLog = "/home/jude/Desktop/amnis_server/log"

func main() {
	confLogger()
	rtsp.InitRtspServer()

	jsonRepo := media.NewJsonRepo()

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

func confLogger() {
	openFile, err := os.Create(outLog)
	if err != nil {
		log.Fatal()
		return
	}
	err = openFile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	devConf := zap.NewProductionConfig()
	devConf.OutputPaths = []string{"stdout", outLog}
	appLogger, err := devConf.Build()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(appLogger)
}
