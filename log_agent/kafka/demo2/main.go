// 通过读配置文件config.ini初始化kafka

package main

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"

	"logagent_study/kafka"
	"logagent_study/taillog"
)

func run() {
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

	//
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed, err:%v\n", err)
		return
	}
	fmt.Printf("init kafka success")

	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("init taillog failed, err:%v\n", err)
		return
	}
	fmt.Printf("init taillog success")

	//
	run()
}
