package {{.PkgName}}

import (
    "content_svr/common/result"
    "github.com/go-playground/validator/v10"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

type verify{{.RequestType}}Request struct {
	types.{{.RequestType}}
}

func (p *verify{{.RequestType}}Request) Validate() error {
	valid := validator.New()
	if err := valid.Struct(p); err != nil {
		return err
	}
	return nil
}


func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req verify{{.RequestType}}Request
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req.{{.RequestType}}{{end}})
		result.HttpResult(r, w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
	}
}
