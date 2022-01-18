package middlerware

import "testing"

func TestJWT(t *testing.T) {
	JWTAuth()
	SetSignKey("x")
}
