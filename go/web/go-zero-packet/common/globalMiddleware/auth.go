package globalMiddleware

import (
	"demo/app/login/rpc/login_client"
	"demo/app/login/rpc/pb/login"
	"demo/common/result"
	"errors"
	"net/http"
	"strconv"
)

type AuthMiddleware struct {
	RpcLogin login_client.Login
}

func NewAuthMiddleware(rpcLogin login_client.Login) *AuthMiddleware {
	return &AuthMiddleware{RpcLogin: rpcLogin}
}

func (p *AuthMiddleware) Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			result.ForbiddenErrorResult(r, w, errors.New("token is empty"))
			return
		}

		userInfo, err := p.RpcLogin.Auth(r.Context(), &login.AuthReq{Token: token})
		if err != nil {
			result.ForbiddenErrorResult(r, w, errors.New("token is empty"))
			return
		}

		r.Header.Set("userId", strconv.Itoa(int(userInfo.UserId)))

		next(w, r)
		return
	}
}

func InterceptorAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("请求未携带token，无权限访问"))
			return
		}

		// todo rpc

		next(w, r)
	}
}
