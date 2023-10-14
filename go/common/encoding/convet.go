package encoding

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"net"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/net/html"
)

// StringToUnicode 中文转unicode编码
func StringToUnicode(value string) (string, string) {
	rs := []rune(value)
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 16 {
			json += "\\u000" + strconv.FormatInt(int64(rint), 16) // json
		} else if rint < 256 {
			json += "\\u00" + strconv.FormatInt(int64(rint), 16) // json
		} else if rint < 4096 {
			json += "\\u0" + strconv.FormatInt(int64(rint), 16) // json
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return json, strings.ReplaceAll(json, `\u`, "%u")
}

// UnicodeToString unicode编码转中文
func UnicodeToString(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		_ = binary.Read(br, binary.BigEndian, &r)
		to += string([]rune{rune(r)})
	}
	return
}

// StrToASCII 字符转ASCII码
func StrToASCII(str string) string {
	rs := []rune(str)
	nums := make([]string, len(rs))
	for i := range rs {
		num := int(rs[i])
		nums[i] = strconv.Itoa(num)
	}
	return strings.Join(nums, ";")
}

// AsciiToStr ASCII码转字符串
func AsciiToStr(str string) (string, error) {
	nums := strings.Split(str, ";")
	rs := make([]rune, 0, len(nums))
	for i := range nums {
		if num, err := strconv.Atoi(nums[i]); err != nil {
			return "", err
		} else {
			rs = append(rs, rune(num))
		}
	}
	return string(rs), nil
}

// AsciiToChar ASCII码转字符
func AsciiToChar(i int) string {
	return string(rune(i))
}

// DecBin 十进制转二进制
func DecBin(n int64) string {
	return strconv.FormatUint(*(*uint64)(unsafe.Pointer(&n)), 2)
}

// DecOct 十进制转八进制
func DecOct(d int64) (int64, error) {
	if d == 0 {
		return 0, nil
	}
	if d < 0 {
		return -1, errors.New("decimal to octal error: the argument must be greater than zero")
	}
	s := ""
	for q := d; q > 0; q = q / 8 {
		m := q % 8
		s = fmt.Sprintf("%v%v", m, s)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return int64(n), nil
}

// DecHex 十进制转十六进制
func DecHex(n int64) (string, error) {
	if n < 0 {
		return "", errors.New("decimal to hexadecimal error: the argument must be greater than zero")
	}
	if n == 0 {
		return "0", nil
	}
	hex_ := map[int64]rune{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			l := hex_[m]
			s = fmt.Sprintf("%v%v", string(l), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s, nil
}

// BinDec 二进制转十进制
func BinDec(b string) (n int64, err error) {
	s := strings.Split(b, "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			return -1, err
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int64(d), nil
}

// OctDec 八进制转十进制
func OctDec(o int64) (n int64, err error) {
	s := strings.Split(strconv.Itoa(int(o)), "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			return -1, err
		}
		d += f * math.Pow(8, float64(l-i-1))
	}
	return int64(d), nil
}

// HexDec 十六进制转十进制
func HexDec(h string) (n int64, err error) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex_ := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex_[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			return -1, err
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d), nil
}

// OctBin 八进制转二进制
func OctBin(o int64) (string, error) {
	d, err := OctDec(o)
	if err != nil {
		return "", err
	}
	return DecBin(d), nil
}

// OctHex 八进制转十六进制
func OctHex(o int64) (string, error) {
	d, err := OctDec(o)
	if err != nil {
		return "", err
	}
	return DecHex(d)
}

// HexOct 十六进制转八进制
func HexOct(h string) (int64, error) {
	d, err := HexDec(h)
	if err != nil {
		return -1, err
	}
	return DecOct(d)
}

// HexBin 十六进制转二进制
func HexBin(h string) (string, error) {
	d, err := HexDec(h)
	if err != nil {
		return "", err
	}
	return DecBin(d), nil
}

// BinOct 二进制转八进制
func BinOct(b string) (int64, error) {
	d, err := BinDec(b)
	if err != nil {
		return -1, err
	}
	return DecOct(d)
}

// BinHex 二进制转十六进制
func BinHex(b string) (string, error) {
	d, err := BinDec(b)
	if err != nil {
		return "", err
	}
	return DecHex(d)
}

// IPString2Long 把ip字符串转为数值
func IPString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}
	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// Long2IPString 把数值转为ip字符串
func Long2IPString(i uint) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}

func EncodeHtml(htmlStr string) string {
	return html.EscapeString(htmlStr)
}

func DecodeHtml(htmlStr string) string {
	return html.UnescapeString(htmlStr)
}

// StringToMssqlStr 字符串转mssql字符串
func StringToMssqlStr(str string) string {
	rs := []rune(str)
	strs := make([]string, len(rs))
	for i := range rs {
		strs[i] = "NCHAR(" + strconv.Itoa(int(rs[i])) + ")"
	}
	return strings.Join(strs, "+")
}

// MssqlStrToString mssql字符串转字符串
func MssqlStrToString(mssqlStr string) string {
	chars := strings.Split(mssqlStr, "+")
	rs, err := dbStrToString(chars, "nchar(")
	if err != nil {
		return mssqlStr
	}
	return string(rs)
}

// StringToOracleStr 字符串转oracle字符串
func StringToOracleStr(str string) string {
	rs := []rune(str)
	strs := make([]string, len(rs))
	for i := range rs {
		strs[i] = "CHR(" + strconv.Itoa(int(rs[i])) + ")"
	}
	return strings.Join(strs, "||")
}

// OracleStrToString oracle字符串转字符串
func OracleStrToString(oracleStr string) string {
	chars := strings.Split(oracleStr, "||")
	rs, err := dbStrToString(chars, "chr(")
	if err != nil {
		return oracleStr
	}
	return string(rs)
}

func dbStrToString(chars []string, chr string) ([]rune, error) {
	rs := make([]rune, 0)
	for _, char := range chars {
		if len(char) > len(chr)+1 {
			index := strings.Index(strings.ToLower(char), chr)
			if index < 0 {
				goto loop
			} else if index == 0 {
				cr := char[index+len(chr) : strings.Index(char, ")")]
				num, err := strconv.Atoi(cr)
				if err == nil {
					rs = append(rs, rune(num))
					index += len(cr) + len(chr)
				} else {
					return nil, err
				}
			} else {
				for _, c := range char[:index] {
					rs = append(rs, c)
				}
				cr := char[index+len(chr) : strings.Index(char, ")")]
				num, err := strconv.Atoi(cr)
				if err == nil {
					rs = append(rs, rune(num))
					index += len(cr) + len(chr)
				} else {
					return nil, err
				}
			}
			for _, c := range char[index+1:] {
				rs = append(rs, c)
			}
			continue
		}
	loop:
		for _, c := range char {
			rs = append(rs, c)
		}
	}
	return rs, nil
}

// StringFromCharCode String.fromCharCode
func StringFromCharCode(str string) string {
	rs := []rune(str)
	strs := make([]string, len(rs))
	for i := range rs {
		strs[i] = strconv.Itoa(int(rs[i]))
	}
	str = strings.Join(strs, ",")
	return "String.fromCharCode(" + str + ")"
}

func UnStringFromCharCode(str string) (string, error) {
	index1 := strings.LastIndex(str, "(")
	index2 := strings.LastIndex(str, ")")
	str = str[index1+1 : index2]
	strs := strings.Split(str, ",")
	rs := make([]rune, 0, len(strs))
	for i := range strs {
		if strs[i] != "" {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				return "", err
			}
			rs = append(rs, rune(num))
		}
	}
	return string(rs), nil
}

// StcFieldTagMap 获取struct字段名与tag的映射关系，并存储到tf中(map[tag]name)
func StcFieldTagMap(v reflect.Value, tagName string, tf map[string]string) {
	stcTag := make([]string, 2)
	for i := 0; i < v.NumField(); i++ {
		t := v.Type().Field(i)
		tag, ok := t.Tag.Lookup(tagName)
		if !ok || tag == "-" {
			continue
		}
		kind := t.Type.Kind()
		field := v.Field(i)
		if kind == reflect.Ptr {
			field = field.Elem()
		}
		if field.Kind() == reflect.Struct {
			stcTag[0] = tag
			StcFieldTagMap(field, tagName, tf)
			continue
		}
		if field.Kind() == reflect.Interface {
			wrapped := field.Elem()
			if wrapped.Kind() == reflect.Ptr {
				wrapped = wrapped.Elem()
			}
			StcFieldTagMap(wrapped, tagName, tf)
		}
		tf[tag] = t.Name
	}
	return
}

// Interface2String interface转string
func Interface2String(i interface{}) string {
	value := reflect.ValueOf(i)
	switch value.Kind() {
	case reflect.String:
		return i.(string)
	case reflect.Bool:
		return strconv.FormatBool(i.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%d", i)
	case reflect.Float32:
		return strconv.FormatFloat(float64(i.(float32)), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case reflect.Complex64:
		return strconv.FormatComplex(complex128(i.(complex64)), 'g', -1, 64)
	case reflect.Complex128:
		return strconv.FormatComplex(i.(complex128), 'g', -1, 128)
	case reflect.Slice:
		if b, ok := i.([]byte); ok {
			return string(b)
		}
		if s, ok := i.([]string); ok {
			return strings.Join(s, "")
		}
		return fmt.Sprintf("%v", i)
	case reflect.Ptr:
		if !value.IsNil() {
			return Interface2String(value.Elem().Interface())
		}
		return ""
	}
	return ""
}
