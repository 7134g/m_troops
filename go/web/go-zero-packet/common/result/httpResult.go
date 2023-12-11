package result

import (
	"demo/common/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errCode, errMessage := xerr.Switch(err)
		logx.WithContext(r.Context()).Errorf("Error: %+v ", err)
		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMessage))
	}

}

// ParamErrorResult http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	logx.WithContext(r.Context()).Errorf("Params Error: %+v ", err)
	httpx.WriteJson(w, http.StatusOK, Error(xerr.SvcCodeParams, xerr.SvcMsgParamError))
}

func ForbiddenErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	logx.WithContext(r.Context()).Errorf("Forbidden Error: %+v ", err)
	httpx.WriteJson(w, http.StatusOK, Error(xerr.SvcCodeForbidden, xerr.SvcMsgForbidden))
}
