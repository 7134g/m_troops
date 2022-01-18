package encode

import (
	"encoding/base32"
	"math/rand"
	"time"
)

// GenRandomString 随机生成字符串
func GenRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result []byte
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
const StdPadding = 'w'

var (
	encode = base32.NewEncoding(encodeStd).WithPadding(StdPadding)
)

// EncodeAZ 编码
func EncodeAZ(s string) string {
	return encode.EncodeToString([]byte(s))
}

// DecodeAZ 解码
func DecodeAZ(s string) (string, error) {
	d, err := encode.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
