package encoding

import (
	"reflect"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	type args struct {
		crypt []byte
		key   []byte
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
			if got := AESDecrypt(tt.args.crypt, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AESDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAESEncrypt(t *testing.T) {
	type args struct {
		origData []byte
		key      []byte
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
			if got := AESEncrypt(tt.args.origData, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AESEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPKCS7Padding(t *testing.T) {
	type args struct {
		origData  []byte
		blockSize int
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
			if got := PKCS7Padding(tt.args.origData, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS7Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPKCS7UnPadding(t *testing.T) {
	type args struct {
		origData []byte
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
			if got := PKCS7UnPadding(tt.args.origData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS7UnPadding() = %v, want %v", got, tt.want)
			}
		})
	}
}
