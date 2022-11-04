package main

import (
	"fmt"
	"testing"
)

func TestCompression(t *testing.T) {
	testData := []byte("aaaaaaaaaaa")
	fmt.Println(testData)
	pdata := Compression(testData)
	fmt.Println(pdata)
	result := Decompress(pdata)
	fmt.Println(result)
}
