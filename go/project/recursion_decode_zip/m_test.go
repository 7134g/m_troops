package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestNewUnzip(t *testing.T) {
	uz, err := NewDeCompressZip("nothing.zip", "./")
	if err != nil {
		log.Fatal(err)
	}

	passChan := make(chan string, 10)
	go func() {
		for i := 111111; i < 1000000; i++ {
			passChan <- fmt.Sprintf("%d", i)
		}
	}()

	time.Sleep(time.Second)
	uz.SetPasswdTask(passChan, 3)
	if err := uz.run(); err != nil {
		log.Println(err)
	}

	//if _, err := uz.tryDecrypt("111111"); err != nil {
	//	log.Println(err)
	//}
	//
	//if _, err := uz.tryDecrypt("123456"); err != nil {
	//	log.Println(err)
	//}
}
