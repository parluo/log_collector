/*
undefined: resolver.BuildOption、undefined: balancer.PickOptions
ETCD中使用的旧版本gRPC库与最新版本的gRPC库不兼容，需要在go.mod中将gRPC替换为v1.26.0版本的即可。
*/
package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "logagent/conf", `[
		{
			"path":"./conf/test/web.log", 
			"topic":"web_log"
		},
		{
			"path":"./conf/test/redis.log",
			"topic":"redis_log"
		}
	]`)
	// _, err = cli.KV.Put(ctx, "qimi", "nimaSB")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "logagent/conf")
	// resp, err := cli.KV.Get(ctx, "qimi")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
