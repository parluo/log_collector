package main

import (
	"fmt"
	"kafka_study/conf"

	"gopkg.in/ini.v1"
)

func main() {
	p := new(conf.AppConf)
	ini.MapTo(&p, "conf/config.ini")

	fmt.Println(p)
}
