package logic

import (
	"context"
	"demo/app/login/rpc/internal/svc"
	"demo/app/login/rpc/pb/login"
	"demo/common/jwt"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginResp, error) {
	if !(l.svcCtx.Config.User.Username == in.Username && l.svcCtx.Config.User.Password == in.Password) {
		return nil, errors.New("login error")
	}
	var userId int64 = 1

	j := jwt.NewJWT()
	token, err := j.CreateToken(jwt.CustomClaims{
		UserId: userId,
		Expire: time.Now().Add(time.Hour * 24 * 7).UnixMilli(),
	})
	if err != nil {
		return nil, err
	}

	rdsKey := fmt.Sprintf(jwt.CacheKey, userId)
	if err := l.svcCtx.RedisClint.Set(rdsKey, token); err != nil {
		logx.Error(err)
	}

	return &login.LoginResp{Token: token}, nil
}
