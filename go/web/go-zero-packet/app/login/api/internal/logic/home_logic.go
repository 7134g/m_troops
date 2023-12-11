package logic

import (
	"context"

	"demo/app/login/api/internal/svc"
	"demo/app/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeLogic {
	return &HomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeLogic) Home(req *types.HomeRequest) (resp *types.HomeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
