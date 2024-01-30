package proxy

import (
	"net/http/httptest"
	"testing"
)

func TestNewWrite(t *testing.T) {
	response := &httptest.ResponseRecorder{}
	lw := newWrite(response)
	_, _ = lw.Write([]byte("abc"))
	t.Log(lw.responseBody.String())
}
