package main

import (
	"net"
	"time"
)

func PortToPort(port1 string, port2 string) {
	listen1 := start_server("0.0.0.0:" + port1)
	listen2 := start_server("0.0.0.0:" + port2)
	log.Println("[√]", "listen port:", port1, "and", port2, "success. waiting for client...")
	for {
		conn1 := accept(listen1)
		conn2 := accept(listen2)
		if conn1 == nil || conn2 == nil {
			log.Println("[x] accept client faild. retry in ", timeout, " seconds. ")
			time.Sleep(timeout * time.Second)
			continue
		}
		forward(conn1, conn2)
	}
}

func PortToHost(allowPort string, targetAddress string) {
	server := start_server("0.0.0.0:" + allowPort)
	for {
		conn := accept(server)
		if conn == nil {
			continue
		}

		log.Println("[+] start connect host:[" + targetAddress + "]")
		for {
			target, err := net.Dial("tcp", targetAddress)
			if err != nil {
				// temporarily unavailable, don't use fatal.
				log.Println("[x] connect target address ["+targetAddress+"] faild. retry in ", timeout, "seconds. ")
				log.Println("[←] close the connect at local:[" + conn.LocalAddr().String() + "] and remote:[" + conn.RemoteAddr().String() + "]")
				time.Sleep(timeout * time.Second)
				log.Println("[+] restart connect[" + targetAddress + "]")
				continue
			}
			log.Println("[→] connect target address [" + targetAddress + "] success.")
			forward(conn, target)
			break
		}

		//go func(targetAddress string) {
		//	log.Println("[+]", "start connect host:["+targetAddress+"]")
		//	target, err := net.Dial("tcp", targetAddress)
		//	if err != nil {
		//		// temporarily unavailable, don't use fatal.
		//		log.Println("[x] connect target address ["+targetAddress+"] faild. retry in ", timeout, "seconds. ")
		//		//conn.Close()
		//		log.Println("[←]", "close the connect at local:["+conn.LocalAddr().String()+"] and remote:["+conn.RemoteAddr().String()+"]")
		//		//time.Sleep(timeout * time.Second)
		//		return
		//	}
		//	log.Println("[→]", "connect target address ["+targetAddress+"] success.")
		//	forward(conn, target)
		//}(targetAddress)
	}
}

func HostToHost(address1, address2 string) {
	for {
		log.Println("[+]", "try to connect host:["+address1+"] and ["+address2+"]")
		var host1, host2 net.Conn
		var err error
		for {
			host1, err = net.Dial("tcp", address1)
			if err == nil {
				log.Println("[→]", "connect ["+address1+"] success.")
				break
			} else {
				log.Println("[x] connect target address ["+address1+"] faild. retry in ", timeout, " seconds. ")
				time.Sleep(1 * time.Second)
			}
			log.Println("[→] try to reconnect: ", address1)
		}
		for {
			host2, err = net.Dial("tcp", address2)
			if err == nil {
				log.Println("[→]", "connect ["+address2+"] success.")
				break
			} else {
				log.Println("[x] connect target address ["+address2+"] faild. retry in ", timeout, " seconds. ")
				time.Sleep(1 * time.Second)
			}
			log.Println("[→] try to reconnect: ", address2)
		}
		forward(host1, host2)
		log.Println("[x] now both 1 and 2 have closed the connection")
		log.Printf("[√] %s and %s start reconnect\n", address1, address2)
	}
}
