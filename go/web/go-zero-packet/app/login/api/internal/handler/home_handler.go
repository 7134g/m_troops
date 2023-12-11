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

type verifyHomeRequestRequest struct {
	types.HomeRequest
}

func (p *verifyHomeRequestRequest) Validate() error {
	valid := validator.New()
	if err := valid.Struct(p); err != nil {
		return err
	}
	return nil
}

func HomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req verifyHomeRequestRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewHomeLogic(r.Context(), svcCtx)
		resp, err := l.Home(&req.HomeRequest)
		result.HttpResult(r, w, resp, err)
	}
}
