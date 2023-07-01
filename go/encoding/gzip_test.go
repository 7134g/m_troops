package encoding

import (
	"reflect"
	"testing"
)

func TestCompression(t *testing.T) {
	type args struct {
		oriData []byte
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
			if got := Compression(tt.args.oriData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecompress(t *testing.T) {
	type args struct {
		oriData []byte
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
			if got := Decompress(tt.args.oriData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
