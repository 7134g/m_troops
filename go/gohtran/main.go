package main

import (
	"fmt"
	"os"
	"strings"
)

var log = Logger{status: 1}

func main() {
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	//log.SetFlags(log.Ldate | log.Lmicroseconds)

	//printWelcome()
	//
	////flag.IntVar(&AESSIGN, "aes", 0, "AES encryption is not used by default, with 1 for encryption and 2 for decryption")
	////flag.IntVar(&GZIPSIGN, "gzip", 0, "GZIP compression is not used by default, 1 for compression and 2 for decompression")
	//flag.StringVar(&LOGTPYE, "log", ".", "write log")
	//flag.StringVar(&CONFUSE, "key", "^($*(897@6>8<1?9", "The default is used if the key is not specified, len(key) > 15")
	//
	//flag.Var(NewSliceValue([]string{}, &SIGN), "sign", `-listen need two arguments, like "gh -listen 8001,8002".`)
	//flag.Var(NewSliceValue([]string{}, &LISTEN), "listen", `-listen need two arguments, like "gh -listen 8001,8002".`)
	//flag.Var(NewSliceValue([]string{}, &TRAN), "tran", `-tran need two arguments, like "gh -tran 1997,192.168.1.2:3389".`)
	//flag.Var(NewSliceValue([]string{}, &SLAVE), "slave", `-slave need two arguments, like "gh -slave 127.0.0.1:3389,8.8.8.8:1997".`)
	//flag.Parse()

	args := os.Args
	argc := len(os.Args)

	if argc == 1 {
		log.Fatalln("[x] Insufficient incoming parameters")
	}

	// 参数解析
	design := 0
	for i := 0; i < argc; i++ {
		switch args[i] {
		case "-silent":
			log.status = 0
		case "-s":
			log.status = 0
		case "-h", "-help":
			printHelp()
			return
		case "-v", "-version":
			fmt.Println("版 本 号:v1.0")
			fmt.Println("适用平台:Windows/Linux x86/x64")
			fmt.Println("所需权限:无")
			fmt.Println("外部依赖:无")
			return
		}
	}

	for i := 1; i < argc; i++ {
		if ok := strings.Contains(args[i], "-"); ok {
			// other args
			switch args[i] {
			case "-e":
				log.Println("[√] enable AES,GZIP functionality")
				SIGN = append(SIGN, "default")
				checkKey(&i, args)
				continue
			case "-aes":
				log.Println("[√] enable AES functionality")
				SIGN = append(SIGN, "aes")
				checkKey(&i, args)
				continue
			case "-gzip":
				log.Println("[√] enable GZIP functionality")
				SIGN = append(SIGN, "gzip")
				continue
			case "-log":
				log.Println("[√] enable LOG functionality")
				LOGTPYE = "."
				log.SetLocalFile()
				continue

			}

			// design args insufficient
			if argc <= i+2 {
				log.Fatalln("[x] modele is missing a parameter or parameter error")
			}

			// 确保只操作一次 design
			if design == 0 {
				design++
			} else {
				log.Fatalln("[x] Both modes are passed in or parameter error")
			}

			//fmt.Println(args[i])
			switch args[i] {
			case "-listen":
				i++
				LISTEN = append(LISTEN, args[i])
				i++
				LISTEN = append(LISTEN, args[i])
			case "-tran":
				i++
				TRAN = append(TRAN, args[i])
				i++
				TRAN = append(TRAN, args[i])
			case "-slave":
				i++
				SLAVE = append(SLAVE, args[i])
				i++
				SLAVE = append(SLAVE, args[i])
			}

		} else {
			//fmt.Println(args[i], design, LISTEN)
			log.Fatalln("[x] Passing in parameters does not comply with the specification")
		}
	}

	if len(SIGN) == 2 {
		SIGN = []string{"default"}
	}

	//fmt.Println("------")
	//fmt.Println(SIGN)
	//fmt.Println(CONFUSE)
	//fmt.Println(LOGTPYE)
	//fmt.Println(LISTEN)
	//fmt.Println(TRAN)
	//fmt.Println(SLAVE)
	//fmt.Println("------")

	// reload args
	if len(LISTEN) == 0 && len(TRAN) == 0 && len(SLAVE) == 0 {
		log.Fatalln("[x] you should choose either -listen ,-tran or -slave")
	}

	if len(LISTEN) != 0 {
		port1 := checkPort(LISTEN[0])
		port2 := checkPort(LISTEN[1])
		log.Println("[√] start to listen port:", port1, "and port:", port2)
		PortToPort(port1, port2)
		return
	}

	if len(TRAN) != 0 {
		port := checkPort(TRAN[0])
		var remoteAddress string
		if checkIp(TRAN[1]) {
			remoteAddress = TRAN[1]
		}
		//split := strings.SplitN(remoteAddress, ":", 2)
		//checkPortExist(remoteAddress)
		log.Println("[√]", "start to transmit address: 127.0.0.1:"+port, "to address:", remoteAddress)
		PortToHost(port, remoteAddress)
		return
	}

	if len(SLAVE) != 0 {
		var address1, address2 string
		if checkIp(SLAVE[0]) {
			address1 = SLAVE[0]
		}
		if checkIp(SLAVE[1]) {
			address2 = SLAVE[1]
		}
		log.Println("[√]", "start to connect address:", address1, "and address:", address2)
		HostToHost(address1, address2)
		return
	}

}

func printHelp() {
	fmt.Println("+-----------------------------help information--------------------------------+")
	fmt.Println(`usage: "-listen port1 port2" #example: "gohtran -listen 8888 3389" `)
	fmt.Println(`       "-tran port1 ip:port2" #example: "gohtran -tran 8888 1.1.1.1:3389" `)
	fmt.Println(`       "-slave ip1:port1 ip2:port2" #example: "gohtran -slave 127.0.0.1:3389 1.1.1.1:8888" `)
	fmt.Println(`       "-e enable gzip and aes functionality`)
	fmt.Println(`       "-aes enable aes functionality, parameters is key, defaults to 16 bits`)
	fmt.Println(`       "-gzip enable gzip functionality`)
	fmt.Println(`       "-h -help program documentation`)
	fmt.Println(`       "-s -silent silent mode,no information is displayed`)
	fmt.Println(`       "-log output transferred data to file`)
	fmt.Println(`============================================================`)
	fmt.Println("If you see start transmit, that means the data channel is established")
	//fmt.Println(`optional argument: "-log logpath" . example: "nb -listen 1997 2017 -log d:/nb" `)
	//fmt.Println(`log filename format: Y_m_d_H_i_s-agrs1-args2-args3.log`)
	//fmt.Println(`============================================================`)
	//fmt.Println(`if you want more help, please read "README.md". `)
}
