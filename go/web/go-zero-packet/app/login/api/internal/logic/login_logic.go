package logic

import (
	"context"
	"demo/app/login/api/internal/svc"
	"demo/app/login/api/internal/types"
	"demo/app/login/rpc/pb/login"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user := &login.LoginReq{}
	_ = copier.Copy(user, req)
	userInfo, err := l.svcCtx.RpcLogin.Login(l.ctx, user)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResponse{}
	resp.Token = userInfo.Token

	return resp, nil
}
