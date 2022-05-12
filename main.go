package main

import "fmt"

func main() {
	getNalHeader(4)
	//initVideoTransmitter()
	a := FUa{}

	fmt.Printf("%v\n", a.serialize([]byte{3, 4, 5, 6, 7, 8, 9}, 3))
}
