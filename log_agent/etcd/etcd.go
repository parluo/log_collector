package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return fmt.Errorf("etcd connect failed")
	}
	return nil
}

func GetConf(key string) (confs []*LogEntry, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	resp, err := cli.Get(ctx, key)

	cancel()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &confs) // 字符数组序列化
		if err != nil {
			fmt.Printf("unmarshal etcd failed, err:%v\n", err)
		}
	}
	return confs, nil
}

func WatchConf(key string, newChanConf chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)

			var NewConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(evt.Kv.Value, &NewConf)
				if err != nil {
					fmt.Printf("unmashal failed, err:%v\n", err)
					continue
				}
			}
			fmt.Printf("get new conf:%v\n", NewConf)
			newChanConf <- NewConf
		}
	}
}
