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

type verifyListReqRequest struct {
	types.ListReq
}

func (p *verifyListReqRequest) Validate() error {
	valid := validator.New()
	if err := valid.Struct(p); err != nil {
		return err
	}
	return nil
}

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req verifyListReqRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req.ListReq)
		result.HttpResult(r, w, resp, err)
	}
}
