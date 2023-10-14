package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Rect struct {
}

type Params struct {
	Width, Height int
}

func (r *Rect) Area(p Params, result *int) error {
	*result = p.Height * p.Width
	return nil
}

func (r *Rect) Perimeter(p Params, result *int) error {
	*result = (p.Width + p.Height) * 2
	return nil
}

func main() {
	rp := new(Rect)
	err := rpc.Register(rp)
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
