package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	args := os.Args

	switch args[1] {
	case "s":
		serve(args[2])
	case "d":
		dial(args[2])
	}
}

var serveConnList = map[string]net.Conn{} // 存放所有dial连接
func serve(address string) {
	defer func() {
		for _, conn := range serveConnList {
			_ = conn.Close()
		}
	}()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		log.Println("accept...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("get new connect %p\n", conn)

		key := fmt.Sprintf("%p", conn)
		serveConnList[key] = conn
		go read(conn)
		go write(conn)
	}
}

func dial(address string) {
	var wg sync.WaitGroup
	wg.Add(2)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("dial success connenct")
	go func() {
		defer wg.Done()
		read(conn)
	}()
	go func() {
		defer wg.Done()
		write(conn)
	}()
	wg.Wait()
}

func read(conn net.Conn) {
	var buff = make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("eof", string(buff[:n]))
				return
			}
			log.Println(err)
			return
		}

		if n == 0 {
			continue
		}

		fmt.Printf("%d -> %s\n", n, string(buff[:n]))
	}
}

func write(conn net.Conn) {
	for {
		var text string
		_, err := fmt.Scan(&text)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("eof", text)
				return
			}
			log.Println(err)
			return
		}
	}
}
