package handler

import (
	"demo/common/result"
	"github.com/go-playground/validator/v10"
	"net/http"

	"demo/app/login/api/internal/logic"
	"demo/app/login/api/internal/svc"
	"demo/app/login/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type verifyLoginRequestRequest struct {
	types.LoginRequest
}

func (p *verifyLoginRequestRequest) Validate() error {
	valid := validator.New()
	if err := valid.Struct(p); err != nil {
		return err
	}
	return nil
}

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req verifyLoginRequestRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req.LoginRequest)
		result.HttpResult(r, w, resp, err)
	}
}
