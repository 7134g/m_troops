package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

func main() {
	rp, err := jsonrpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	err = rp.Call("Rect.Area", Params{50, 100}, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("area: ", result)

	err = rp.Call("Rect.Perimeter", Params{50, 100}, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("perometer: ", result)

}
