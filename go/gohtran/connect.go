package main

import (
	"io"
	"net"
	"sync"
)

func start_server(address string) net.Listener {
	log.Println("[+]", "try to start server on:["+address+"]")
	server, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("[x] listen address [" + address + "] faild.")
	}
	log.Println("[√]", "start listen at address:["+address+"]")
	return server
}

func accept(listener net.Listener) net.Conn {
	conn, err := listener.Accept()
	if conn == nil {
		log.Println("[x] connect is nil")
		return nil
	}
	if err != nil {
		log.Println("[x] accept connect ["+conn.RemoteAddr().String()+"] faild.", err.Error())
		return nil
	}
	log.Println("[√]", "accept a new client. remote address:["+conn.RemoteAddr().String()+"], local address:["+conn.LocalAddr().String()+"]")
	return conn
}

func forward(conn1 net.Conn, conn2 net.Conn) {
	log.Printf("[+] start transmit. [%s] <-> [%s] and [%s] <-> [%s] \n", conn1.LocalAddr().String(), conn1.RemoteAddr().String(), conn2.LocalAddr().String(), conn2.RemoteAddr().String())
	var wg sync.WaitGroup
	// wait tow goroutines
	wg.Add(2)
	go connCopy(conn1, conn2, &wg, "lr")
	go connCopy(conn2, conn1, &wg, "rl")
	//blocking when the wg is locked
	wg.Wait()
}

func connCopy(conn1 net.Conn, conn2 net.Conn, wg *sync.WaitGroup, flag string) {
	defer conn1.Close()
	if len(SIGN) == 0 {
		//logFile := openLog(conn1.LocalAddr().String(), conn1.RemoteAddr().String(), conn2.LocalAddr().String(), conn2.RemoteAddr().String())
		//if logFile != nil {
		//	w := io.MultiWriter(conn1, logFile)
		//	io.Copy(w, conn2)
		//} else {
		//	io.Copy(conn1, conn2)
		//}
		io.Copy(conn1, conn2)
		//conn2.Close()
		//log.Println("[←]", "close the connect at local:["+conn2.LocalAddr().String()+"] and remote:["+conn2.RemoteAddr().String()+"]")
	} else {
		ConnRW(conn1, conn2)
	}
	log.Println("[←]", "close the connect at local:["+conn1.LocalAddr().String()+"] and remote:["+conn1.RemoteAddr().String()+"]")
	wg.Done()
}
