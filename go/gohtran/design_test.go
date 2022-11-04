package main

import (
	"testing"
)

func TestPortToPort(t *testing.T) {
	//AESSIGN = 1
	//CONFUSE = "^($*(897@6>8<1?9"
	//SIGNCHAN <- 3
	//LOGTPYE = "."
	//var wg sync.WaitGroup
	//wg.Add(4)
	SIGN = []string{"aes"}
	PortToPort("2000", "2100")
}

func TestPortToHost(t *testing.T) {
	//LOGTPYE = "."
	//AESSIGN = 1
	//CONFUSE = "^($*(897@6>8<1?9"
	//SIGNCHAN <- 3
	SIGN = []string{"aes"}
	port := "1800"
	//address := "127.0.0.1:2000"
	address := "192.168.202.128:2000"
	checkPort(port)
	checkIp(address)
	//checkPortExist(address)
	PortToHost(port, address)
}

func TestHostToHost(t *testing.T) {
	HostToHost("127.0.0.1:3000", "127.0.0.1:3100")
}
