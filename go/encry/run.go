package main

import (
	"encry/config"
	"fmt"
	"log"
	"strings"
)

func Run() {

}

func loadOS(args []string, argLen int) {
	for i := 1; i < argLen; i++ {
		if ok := strings.Contains(args[i], "-"); ok {
			// other args
			switch args[i] {
			case "-l":
				i++
				config.LOCALPORT = args[i]
			case "-local":
				i++
				config.LOCALPORT = args[i]
			case "-r":
				i++
				config.REMOTEADDRESS = args[i]
			case "-remote":
				i++
				config.REMOTEADDRESS = args[i]
			case "-m", "-mute": // 静默
				config.LOGSILENT = 1
			case "-c", "-client":
				config.StartServer = 1
			case "-s", "-server":
				config.StartClient = 1
			case "-h":
				help()
			}
		}
	}

}

func loadYaml() {
	config.LoadYaml()
}

func help() {
	fmt.Println("+-----------------------------help information--------------------------------+")
	fmt.Println(`usage: "-listen port1 port2" #example: "" `)
	fmt.Println(`============================================================`)
	fmt.Println("If you see xxxxxx, that means the data channel is established")
}

func welcome() {
	log.Println("============== welcome ==================")
	log.Println("Program execution begins...")
}
