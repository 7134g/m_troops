package model

import (
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MonConfig interface {
	GetURL() string
	GetDBName() string
}

var maxPoolSizeOption mon.Option = func(opts *options.ClientOptions) {
	opts.SetMaxPoolSize(100)
}

var timeoutOption mon.Option = func(opts *options.ClientOptions) {
	opts.SetTimeout(time.Minute)
}
