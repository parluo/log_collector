package main

import (
	"fmt"
	"log_transfer/conf"
	"log_transfer/es"
	"log_transfer/kafka"

	"gopkg.in/ini.v1"
)

func main() {
	//加载配置文件
	cfgData := conf.LogTransfer{}
	err := ini.MapTo(&cfgData, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config failed, err%v\n", err)
		return
	}
	fmt.Printf("cfg: %v\n", cfgData)
	//初始化es
	err = es.Init(cfgData.ESCfg.Address)
	if err != nil {
		fmt.Printf("init es client failed: %v", err)
		return
	}
	fmt.Println("init es success...")
	//初始化kafka
	err = kafka.Init([]string{cfgData.Kafka.Address}, cfgData.Kafka.Topic)
	if err != nil {
		fmt.Println("kafka init failed: ", err.Error())
		return
	}
	//从kafka取日志数据
	select {}
}
