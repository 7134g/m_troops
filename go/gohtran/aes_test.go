package main

import (
	"fmt"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	//设置钥匙
	key := []byte("^($*(897@6>8<1?9")
	//明文
	origData := []byte("123456")
	fmt.Println(len(origData), string(origData))

	//加密
	//origData := GetCtx(conn)
	en := AESEncrypt(origData, key)
	fmt.Println(len(en), string(en))
	s := string(en)
	d := []byte(s)

	//解密
	de := AESDecrypt(d, key)
	fmt.Println(len(de), string(de))

}
