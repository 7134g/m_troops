package logic

import (
	"context"
	"demo/app/login/rpc/internal/svc"
	"demo/app/login/rpc/pb/login"
	"demo/common/jwt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *login.AuthReq) (*login.AuthResp, error) {
	resp := &login.AuthResp{Token: in.Token}
	j := jwt.NewJWT()
	claims, err := j.ParseToken(in.Token)
	if err != nil {
		return nil, err
	}

	resp.Verification = true
	resp.UserId = claims.UserId
	now := time.Now()
	tokenTime := time.UnixMilli(claims.Expire)
	if now.After(tokenTime) {
		token, err := j.RefreshToken(in.Token)
		if err != nil {
			return nil, err
		}
		resp.Token = token
	}

	return resp, err
}
