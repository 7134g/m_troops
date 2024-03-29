// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Username string `json:"username" validate:"min=0,max=10"`
	Password string `json:"password" validate:"min=0,max=10"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type HomeRequest struct {
}

type HomeResponse struct {
}
