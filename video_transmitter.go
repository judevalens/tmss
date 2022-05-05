package main

// #include "video_reader.h"
import "C"
import (
	"fmt"
	"net"
)

const receiverAdd = ""
const addr = ":4874"
const mtu = 65527

func transmitH264() {
	var err error
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Printf("failed to create udp connection\nerr:%v", err)
	}
	buff := make([]byte, 65527)
	for {
		read, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("failed to read from udp conn\nerr:%v", err)
			return
		}
		_ = read
		fmt.Printf("receive new ")
	}
}
