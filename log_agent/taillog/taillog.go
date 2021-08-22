package taillog

import (
	"context"
	"fmt"
	"logagent_study/kafka"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail

	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init()
	return //=// return tailObj
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}

	go t.run()
}
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task : %s_%s 已停止\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			fmt.Printf("get log from : %s,  %s\n", t.path, line.Text)
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}
