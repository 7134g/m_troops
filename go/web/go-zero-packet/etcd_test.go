package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:        []string{"192.168.1.26:2379"},
		DialTimeout:      time.Second * 5,
		AutoSyncInterval: time.Second * 5,
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(cli.Put(context.Background(), "test", "123456"))
}
