// Code generated by goctl. DO NOT EDIT.
// Source: login.proto

package server

import (
	"context"

	"demo/app/login/rpc/internal/logic"
	"demo/app/login/rpc/internal/svc"
	"demo/app/login/rpc/pb/login"
)

type LoginServer struct {
	svcCtx *svc.ServiceContext
	login.UnimplementedLoginServer
}

func NewLoginServer(svcCtx *svc.ServiceContext) *LoginServer {
	return &LoginServer{
		svcCtx: svcCtx,
	}
}

func (s *LoginServer) Auth(ctx context.Context, in *login.AuthReq) (*login.AuthResp, error) {
	l := logic.NewAuthLogic(ctx, s.svcCtx)
	return l.Auth(in)
}

func (s *LoginServer) Login(ctx context.Context, in *login.LoginReq) (*login.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
