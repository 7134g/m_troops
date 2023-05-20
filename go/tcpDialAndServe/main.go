package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	args := os.Args

	switch args[1] {
	case "s":
		serve(args[2])
	case "d":
		dial(args[2])
	}
}

func serve(address string) {
	wg.Add(2)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go read(conn)
		go write(conn)
	}
	wg.Wait()
}

func dial(address string) {
	wg.Add(2)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	go read(conn)
	go write(conn)
	wg.Wait()
}

func read(conn net.Conn) {
	defer wg.Done()

	var buff = make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Fatalln(err)
		}

		if n == 0 {
			continue
		}

		fmt.Printf("%d -> %s\n", n, string(buff))
	}
}

func write(conn net.Conn) {
	defer wg.Done()
	for {
		var text string
		_, err := fmt.Scan(&text)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
