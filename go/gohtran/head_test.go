package main

import (
	"fmt"
	"testing"
)

func TestCreateHead(t *testing.T) {
	head := []byte(HEADER)
	//data := []byte("123456")
	//fmt.Println(checkHeader(head))
	//fmt.Println(string(changeHeadLenth(head, data)))
	data := append(head, []byte("abcd")...)
	data = append(data, head...)
	data = append(data, []byte("abcd")...)
	data = append(data, head...)
	data = append(data, []byte("abcd")...)

	data = changeHeadLenth(data)
	fmt.Println(data[:HEADERLEN])

	indexs := checkHeaderCount(data)
	fmt.Println(indexs)

	length := getHeadLenth(data[CSTART:CEND])
	fmt.Println(length)
	//indexs := checkHeaderCount(data)
	//cop := [][]byte{}
	//
	//for i := 0; i < len(indexs); i++ {
	//	if i == 0 {
	//		cop = append(cop, data[:indexs[i]])
	//	}
	//	if i+1 < len(indexs) {
	//		cop = append(cop, data[indexs[i]:indexs[i+1]])
	//	} else {
	//		cop = append(cop, data[indexs[i]:])
	//	}
	//}
	//
	//fmt.Println(data)
	//fmt.Println(string(cop[0]), string(cop[1]), string(cop[2]))
	//fmt.Println(string(bytes.Join(cop, []byte{})))
}
