package encoding

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

func TestAsciiToChar(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsciiToChar(tt.args.i); got != tt.want {
				t.Errorf("AsciiToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsciiToStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AsciiToStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("AsciiToStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsciiToStr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinDec(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name    string
		args    args
		wantN   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := BinDec(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinDec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("BinDec() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestBinHex(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinHex(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BinHex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinOct(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinOct(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinOct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BinOct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecBin(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecBin(tt.args.n); got != tt.want {
				t.Errorf("DecBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecHex(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecHex(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecHex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecOct(t *testing.T) {
	type args struct {
		d int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecOct(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecOct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecOct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeHtml(t *testing.T) {
	type args struct {
		htmlStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeHtml(tt.args.htmlStr); got != tt.want {
				t.Errorf("DecodeHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeHtml(t *testing.T) {
	type args struct {
		htmlStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeHtml(tt.args.htmlStr); got != tt.want {
				t.Errorf("EncodeHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexBin(t *testing.T) {
	type args struct {
		h string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexBin(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexBin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HexBin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexDec(t *testing.T) {
	type args struct {
		h string
	}
	tests := []struct {
		name    string
		args    args
		wantN   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := HexDec(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexDec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("HexDec() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestHexOct(t *testing.T) {
	type args struct {
		h string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexOct(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexOct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HexOct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPString2Long(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IPString2Long(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPString2Long() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IPString2Long() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLong2IPString(t *testing.T) {
	type args struct {
		i uint
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Long2IPString(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Long2IPString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Long2IPString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctBin(t *testing.T) {
	type args struct {
		o int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OctBin(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("OctBin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OctBin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctDec(t *testing.T) {
	type args struct {
		o int64
	}
	tests := []struct {
		name    string
		args    args
		wantN   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := OctDec(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("OctDec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("OctDec() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestOctHex(t *testing.T) {
	type args struct {
		o int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OctHex(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("OctHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OctHex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToASCII(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToASCII(tt.args.str); got != tt.want {
				t.Errorf("StrToASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringFromCharCode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFromCharCode(tt.args.str); got != tt.want {
				t.Errorf("StringFromCharCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUnicode(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StringToUnicode(tt.args.value)
			if got != tt.want {
				t.Errorf("StringToUnicode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StringToUnicode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUnStringFromCharCode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnStringFromCharCode(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnStringFromCharCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnStringFromCharCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnicodeToString(t *testing.T) {
	type args struct {
		form string
	}
	tests := []struct {
		name    string
		args    args
		wantTo  string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTo, err := UnicodeToString(tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnicodeToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTo != tt.wantTo {
				t.Errorf("UnicodeToString() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}

func Test_dbStrToString(t *testing.T) {
	type args struct {
		chars []string
		chr   string
	}
	tests := []struct {
		name    string
		args    args
		want    []rune
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbStrToString(tt.args.chars, tt.args.chr)
			if (err != nil) != tt.wantErr {
				t.Errorf("dbStrToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbStrToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
