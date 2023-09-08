package main

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	err := DeCompressZip("nothing1.zip", "./nothing", "123456", nil, 0)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func TestDecrypt(t *testing.T) {
	f, err := os.Open("nothing.zip")
	if err != nil {
		t.Fatal(err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	b := make([]byte, info.Size())
	_, _ = f.Read(b)

	pwList := []string{
		"123",
		"111",
	}

	for _, pw := range pwList {
		z := NewZipCrypto([]byte(pw))
		m := z.Decrypt(b)
		fmt.Println(string(m))
		fmt.Println("=======================")
	}

}
