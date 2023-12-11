package handler

import (
	"demo/common/result"
	"github.com/go-playground/validator/v10"
	"net/http"

	"demo/app/meeting/api/internal/logic"
	"demo/app/meeting/api/internal/svc"
	"demo/app/meeting/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type verifyCreateReqRequest struct {
	types.CreateReq
}

func (p *verifyCreateReqRequest) Validate() error {
	valid := validator.New()
	if err := valid.Struct(p); err != nil {
		return err
	}
	return nil
}

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req verifyCreateReqRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req.CreateReq)
		result.HttpResult(r, w, resp, err)
	}
}
