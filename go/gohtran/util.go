package main

import (
	"net"
	"strconv"
	"strings"
)

//func judgeAddress(address string) int {
//	switch address[:3] {
//	case "172":
//		return 1
//	case "192":
//		return 1
//	case "127":
//		return 1
//	case "10":
//		return 1
//	default:
//		return 0
//	}
//}
func checkPortExist(addr string) {
	conn, _ := net.Dial("tcp", addr)
	if conn == nil {
		return
	}
	defer conn.Close()

	// 占用
	log.Fatalln("[x] The connection port is in conflict, port:", addr)
}

func checkPort(port string) string {
	PortNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("[x] port should be a number")
	}
	if PortNum < 1 || PortNum > 65535 {
		log.Fatalln("[x] port should be a number and the range is [1,65536)")
	}
	addr := "127.0.0.1:" + port
	checkPortExist(addr)

	return port
}

func checkIp(address string) bool {
	_, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Fatalln("[x] Incorrect incoming address: " + address)
	}

	return true
}

func checkKey(index *int, args []string) {

	if *index+1 < len(args) {
		if exist := strings.Contains(args[*index+1], "-"); !exist {
			*index++
			//CONFUSE = checkKey(args[index])
			lkey := len(args[*index])
			if lkey != 16 {
				var patch []string
				if lkey < 16 {
					for i := 0; i < 16-lkey; i++ {
						patch = append(patch, strconv.Itoa(0))
					}
					CONFUSE = args[*index] + strings.Join(patch, "")
				} else {
					CONFUSE = args[*index][:16]
				}
			} else {
				CONFUSE = args[*index]
			}
		}
		log.Println("[√] key set successfully")
	}
	return
}

func errorRecover(msg string) {
	if r := recover(); r != nil {
		log.Fatalln(msg)
	}
}

func checkHeader(header []byte) int {
	for i := 0; i < 12; i++ {

		if i >= 4 && i < 8 {
			if !(48 <= header[i] && header[i] <= 57) {
				return 0
			}
			continue
		}

		if i == IGIP || i == IAES {
			if !(MIN_SCOPE <= header[i] && header[i] <= MAX_SCOPE) {
				return 0
			}
			continue
		}
		if header[i] != HEADER[i] {
			return 0
		}
	}

	return 1
}

func checkHeaderCount(data []byte) []int {
	// 检查头部数量
	dataLen := len(data)
	indexs := []int{}

	for i := HEADERLEN; i < len(data); i++ {
		if i+HEADERLEN > dataLen {
			break
		}
		//s := string(data[i])
		//fmt.Println(s)
		// 检查特殊字符
		if data[i] == HEADER[0] && data[i+1] == HEADER[1] && data[i+2] == HEADER[2] && data[i+3] == HEADER[3] &&
			data[i+8] == HEADER[8] && data[i+11] == HEADER[11] {
			head := data[i : i+HEADERLEN]
			// 检查头部
			if checkHeader(head) == 0 {
				i = i + HEADERLEN
			} else {
				indexs = append(indexs, i)
			}
		}
	}
	return indexs
}

func CreateHead() []byte {
	head := []byte(HEADER)
	for _, v := range SIGN {
		switch v {
		case "gzip":
			head[IGIP] = ENCODE
		case "aes":
			head[IAES] = ENCODE
		case "default":
			head[IGIP] = ORIGIN
			head[IAES] = GZIPAES
		}
	}
	return head
}

func deleteHead(head []byte) []byte {
	if head[IGIP] == ORIGIN && head[IAES] == ORIGIN {
		return []byte{}
	}
	return head
}

func getHeadLenth(headLen []byte) int {
	str := string(headLen)
	l, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln("[x] head is error")
	}
	return l
}

func changeHeadLenth(data []byte) []byte {
	//proData := data[HEADERLEN:]
	lenStr := strconv.Itoa(len(data))
	// 补零
	for i := 0; i < 4; i++ {
		if len(lenStr) < 4 {
			lenStr = "0" + lenStr
		}
	}

	// 修改长度
	for i := CSTART; i < CEND; i++ {
		//fmt.Println(string(data[i]))
		data[i] = lenStr[i-CSTART]
	}

	return data
}

func checkAESandGZIP(data []byte) (processData []byte) {
	//fmt.Println(flag)
	var head []byte
	var targetData []byte

	head = data[:HEADERLEN]
	targetData = data[HEADERLEN:]

	// gzip
	switch head[IGIP] {
	case ORIGIN:
		break
	case ENCODE:
		head[IGIP] = DECODE
		targetData = Compression(targetData)
	case DECODE:
		head[IGIP] = ORIGIN
		targetData = Decompress(targetData)
	}

	// aes
	switch head[IAES] {
	case ORIGIN:
		break
	case ENCODE:
		head[IAES] = DECODE
		targetData = AESEncrypt(targetData, []byte(CONFUSE))
	case DECODE:
		head[IAES] = ORIGIN
		targetData = AESDecrypt(targetData, []byte(CONFUSE))
	case GZIPAES:
		head[IAES] = UNGZIPAES
		targetData = Compression(targetData)
		targetData = AESEncrypt(targetData, []byte(CONFUSE))
	case UNGZIPAES:
		head[IAES] = ORIGIN
		targetData = AESDecrypt(targetData, []byte(CONFUSE))
		targetData = Decompress(targetData)
	}

	head = deleteHead(head)

	processData = append(processData, head...)
	processData = append(processData, targetData...)

	return

}
