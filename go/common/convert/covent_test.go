package convert

import (
	"reflect"
	"testing"
)

func TestStcFieldTagMap(t *testing.T) {
	type Addr struct {
		Country  string `json:"country"`
		Province string `json:"province"`
		City     string `json:"city"`
		Info     string `json:"info"`
		Used     *bool  `json:"used"`
		Other    string `json:"-"`
	}
	type Stc struct {
		ID   int `json:"id"`
		User struct {
			Name string  `json:"name"`
			Age  int     `json:"age"`
			Sex  *string `json:"sex"`
		} `json:"user"`
		Addr *Addr `json:"addr"`
	}

	s := &Stc{
		ID: 0,
		User: struct {
			Name string  `json:"name"`
			Age  int     `json:"age"`
			Sex  *string `json:"sex"`
		}{},
		Addr: &Addr{},
	}
	valueOf := reflect.ValueOf(s)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}
	tf := make(map[string]string, valueOf.NumField())
	StcFieldTagMap(valueOf, "json", tf)
	t.Log(tf)
}

func TestStringToMssqlStr(t *testing.T) {
	//b := []rune("123abc567def❤❥웃유")
	//t.Log(b, string(rune(0x266A)))
	t.Log(StringToMssqlStr("123abc567def❤❥웃유"))
}

func TestMssqlStrToString(t *testing.T) {
	arr := []string{
		"abcdefg",
		"abc123def",
		"NCHAR(49)+NCHAR(50)+NCHAR(51)+1+2+3+4+5+6",
		"NCHAR(10084)+NCHAR(10085)+NCHAR(50883)+NCHAR(50976)",
		"abcdefgNCHAR(35)+NCHAR(36)+NCHAR(37)",
		"abc+123+def+NCHAR(35)+NCHAR(36)+NCHAR(37)+NCHAR(+})+()+NCHAR",
		"NCHAR",
		"NCHAR(abc)",
		"NCHAR()",
		"NCHAR(0)",
		"NCHAR(0000)",
		"NCHAR(000050)",
		"NCHAR(97)+NCHAR(98)NCHAR(99)+NCHAR(98)+NCHAR(97)",
		"NCHAR(23041)+NCHAR(23041)",
	}
	for _, item := range arr {
		res := MssqlStrToString(item)
		t.Log(res)
	}
}

func TestStringToOracleStr(t *testing.T) {
	t.Log(StringToOracleStr("123abc567def"))
}

func TestOracleStrToString(t *testing.T) {
	arr := []string{
		"abcdefg",
		"abc123def",
		"CHR(49)||CHR(50)||CHR(51)||1||2||3||4||5||6",
		"CHR(10084)||CHR(10085)||CHR(50883)||CHR(50976)",
		"abcdefgCHR(35)||CHR(36)||CHR(37)",
		"abc||123||def||CHR(35)||CHR(36)||CHR(37)||CHR(||})||()||CHR",
		"CHR",
		"CHR(abc)",
		"CHR()",
		"CHR(0)",
		"CHR(0000)",
		"CHR(000050)",
		"CHR(97)||CHR(98)CHR(99)||CHR(98)||CHR(97)",
	}
	for _, item := range arr {
		res := OracleStrToString(item)
		t.Log(res)
	}
}

func TestInterface2String(t *testing.T) {
	t.Log(Interface2String("123123"))
	t.Log(Interface2String(true))
	t.Log(Interface2String(-123))
	t.Log(Interface2String(int8(-8)))
	t.Log(Interface2String(int16(-16)))
	t.Log(Interface2String(int32(-32)))
	t.Log(Interface2String(int64(-64)))
	t.Log(Interface2String(uint(123)))
	t.Log(Interface2String(uint8(8)))
	t.Log(Interface2String(uint16(16)))
	t.Log(Interface2String(uint32(32)))
	t.Log(Interface2String(uint64(64)))
	t.Log(Interface2String(uintptr(128)))
	t.Log(Interface2String(float32(32.3232323232)))
	t.Log(Interface2String(64.6464644646))
	t.Log(Interface2String(complex64(6 + 4i)))
	t.Log(Interface2String(12 + 8i))
	t.Log(Interface2String([]byte("123")))
	s := true
	t.Log(Interface2String(&s))
}
