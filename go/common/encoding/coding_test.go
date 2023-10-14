package encoding

import (
	"reflect"
	"testing"
)

func TestConvertByte(t *testing.T) {
	type args struct {
		b       []byte
		charset charset
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertByte(tt.args.b, tt.args.charset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	type args struct {
		value string
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
			got, err := DecodeBase64(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeBase64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodePacked(t *testing.T) {
	type args struct {
		value string
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
			got, err := DecodePacked(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodePacked() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodePacked() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeURI(t *testing.T) {
	type args struct {
		value string
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
			got, err := DecodeURI(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeURI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeBase64(t *testing.T) {
	type args struct {
		value string
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
			if got := EncodeBase64(tt.args.value); got != tt.want {
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodePacked(t *testing.T) {
	type args struct {
		value string
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
			got, err := EncodePacked(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodePacked() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodePacked() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeURI(t *testing.T) {
	type args struct {
		value string
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
			if got := EncodeURI(tt.args.value); got != tt.want {
				t.Errorf("EncodeURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEscape(t *testing.T) {
	type args struct {
		value string
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
			got, err := Escape(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Escape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Escape() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGBKToUtf8(t *testing.T) {
	type args struct {
		gbkStr string
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
			got, err := GBKToUtf8(tt.args.gbkStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GBKToUtf8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GBKToUtf8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenRandomString(t *testing.T) {
	type args struct {
		l int
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
			if got := GenRandomString(tt.args.l); got != tt.want {
				t.Errorf("GenRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := HexStr(tt.args.str)
			if got != tt.want {
				t.Errorf("HexStr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HexStr() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("HexStr() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestHtmlEscape(t *testing.T) {
	type args struct {
		htmlstr string
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
			if got := HtmlEscape(tt.args.htmlstr); got != tt.want {
				t.Errorf("HtmlEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHtmlUnescape(t *testing.T) {
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
			got, err := HtmlUnescape(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("HtmlUnescape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HtmlUnescape() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrHex(t *testing.T) {
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
			got, err := StrHex(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrHex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURICompCoding(t *testing.T) {
	type args struct {
		value string
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
			if got := URICompCoding(tt.args.value); got != tt.want {
				t.Errorf("URICompCoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURICompDecoding(t *testing.T) {
	type args struct {
		value string
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
			got, err := URICompDecoding(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("URICompDecoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URICompDecoding() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLEscape(t *testing.T) {
	type args struct {
		path        string
		escapeSlash bool
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
			if got := URLEscape(tt.args.path, tt.args.escapeSlash); got != tt.want {
				t.Errorf("URLEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLUnescape(t *testing.T) {
	type args struct {
		path string
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
			if got := URLUnescape(tt.args.path); got != tt.want {
				t.Errorf("URLUnescape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnescape(t *testing.T) {
	type args struct {
		value string
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
			got, err := Unescape(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unescape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unescape() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtf8ToGBK(t *testing.T) {
	type args struct {
		utf8str string
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
			got, err := Utf8ToGBK(tt.args.utf8str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Utf8ToGBK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Utf8ToGBK() got = %v, want %v", got, tt.want)
			}
		})
	}
}
