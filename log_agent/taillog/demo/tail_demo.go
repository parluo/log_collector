package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main1() {
	filename := "./my.log" // 一直会监听文件末尾，有\n就返回， 删除了，也会回退
	config := tail.Config{  // 这里config 写出从配置文件读
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}

	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text) // 这里怎么控制打印的撤销的？
	}
}
