// 通过读配置文件config.ini初始化kafka

package main

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/ini.v1"

	"logagent_study/conf"
	"logagent_study/etcd"
	"logagent_study/kafka"
	"logagent_study/taillog"
)

var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

func main() {
	// 1 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)

	// 2 初始化kafka连接
	err = kafka.Init(cfg.KafkaConf.Address, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 3 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")

	// //根据ip设独立配置
	// ipAddr, _ := log_util.Getip()
	// etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipAddr)
	logConfs, err := etcd.GetConf(cfg.EtcdConf.Key)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("get conf from etcd: %#v\n", logConfs)

	taillog.Init(logConfs)

	newConfChan := taillog.GetNewConf()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan)
	wg.Wait()

}
