package main

import (
	"context"
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	host := flag.String("d", "127.0.0.1:2379", "address, eq: -host 127.0.0.1:2379")
	mode := flag.String("mode", "", "get|put|get_all")
	key := flag.String("k", "", "key name")
	value := flag.String("v", "", "value")
	flag.Parse()

	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{*host}, // etcd 服务器地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	switch *mode {
	case "get":
		// 读取 key
		resp, err := cli.Get(ctx, *key)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		for _, ev := range resp.Kvs {
			fmt.Printf("%s -> %s\n", ev.Key, ev.Value)
		}
	case "put":
		// 写入 key-value
		_, err = cli.Put(ctx, *key, *value)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	case "get_all":
		// 读取所有 etcdctl get "" --prefix=true
		resp, err := cli.Get(ctx, "", clientv3.WithPrefix())
		if err != nil {
			log.Fatal(err)
		}
		for _, ev := range resp.Kvs {
			fmt.Printf("%s -> %s\n", ev.Key, ev.Value)
		}
	default:
		log.Fatal("no exist mode")
	}

}
