package xlog

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

func LogInit() {
	var cfg logx.LogConf
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"

	logc.MustSetup(cfg)
}
