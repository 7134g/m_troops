package connect

import (
	"encry/config"
	"encry/data_parse"
	"log"
	"strconv"
)

func checkHeader(header []byte) int {
	for i := 0; i < 12; i++ {

		if i >= 4 && i < 8 {
			if !(48 <= header[i] && header[i] <= 57) {
				return 0
			}
			continue
		}

		if i == config.IGIP || i == config.IAES {
			if !(config.MIN_SCOPE <= header[i] && header[i] <= config.MAX_SCOPE) {
				return 0
			}
			continue
		}
		if header[i] != config.HEADER[i] {
			return 0
		}
	}

	return 1
}

func CreateHead() []byte {
	head := []byte(config.HEADER)
	for _, v := range config.SIGN {
		switch v {
		case "gzip":
			head[config.IGIP] = config.ENCODE
		case "aes":
			head[config.IAES] = config.ENCODE
		case "default":
			head[config.IGIP] = config.ORIGIN
			head[config.IAES] = config.GZIPAES
		}
	}
	return head
}

func deleteHead(head []byte) []byte {
	if head[config.IGIP] == config.ORIGIN && head[config.IAES] == config.ORIGIN {
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
	for i := config.CSTART; i < config.CEND; i++ {
		//fmt.Println(string(data[i]))
		data[i] = lenStr[i-config.CSTART]
	}

	return data
}

func checkAESandGZIP(data []byte) (processData []byte) {
	//fmt.Println(flag)
	var head []byte
	var targetData []byte

	head = data[:config.HEADERLEN]
	targetData = data[config.HEADERLEN:]

	// gzip
	switch head[config.IGIP] {
	case config.ORIGIN:
		break
	case config.ENCODE:
		head[config.IGIP] = config.DECODE
		targetData = data_parse.Compression(targetData)
	case config.DECODE:
		head[config.IGIP] = config.ORIGIN
		targetData = data_parse.Decompress(targetData)
	}

	// aes
	switch head[config.IAES] {
	case config.ORIGIN:
		break
	case config.ENCODE:
		head[config.IAES] = config.DECODE
		targetData = data_parse.AESEncrypt(targetData, []byte(config.CONFUSE))
	case config.DECODE:
		head[config.IAES] = config.ORIGIN
		targetData = data_parse.AESDecrypt(targetData, []byte(config.CONFUSE))
	case config.GZIPAES:
		head[config.IAES] = config.UNGZIPAES
		targetData = data_parse.Compression(targetData)
		targetData = data_parse.AESEncrypt(targetData, []byte(config.CONFUSE))
	case config.UNGZIPAES:
		head[config.IAES] = config.ORIGIN
		targetData = data_parse.AESDecrypt(targetData, []byte(config.CONFUSE))
		targetData = data_parse.Decompress(targetData)
	}

	head = deleteHead(head)

	processData = append(processData, head...)
	processData = append(processData, targetData...)

	return

}
