package taillog

import (
	"fmt"
	"logagent_study/etcd"
	"time"
)

var (
	taskMgr *tailLogMgr
)

type tailLogMgr struct {
	logEntryList []*etcd.LogEntry
	taskMap      map[string]*TailTask
	newConfChan  chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntryList: logEntryConf,
		taskMap:      make(map[string]*TailTask, 16),
		newConfChan:  make(chan []*etcd.LogEntry),
	}

	// 初始化时，etcd有多少就起多少个
	for _, logEntry := range logEntryConf {
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		k := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMgr.taskMap[k] = tailObj
	}
	go taskMgr.run()
}

func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan: // 这里是不是只会接受一条？ 写了已有的两个配置却没有，删除两个
			fmt.Printf("tailLogMgr 新配置： %v\n", newConf)
			for _, conf := range newConf {
				k := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := taskMgr.taskMap[k] // 原先mgr中没有的要新增
				if ok {
					continue
				} else {
					fmt.Printf("tailLogMgr 添加了一个配置监听\n")
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.logEntryList = append(t.logEntryList, conf)
					taskMgr.taskMap[k] = tailObj
				}
			}

			// 热加载：文件更新要全部重新读取的
			// 所以， 原先有的要先删除
			for _, c1 := range t.logEntryList {
				isDelete := true
				for _, c2 := range newConf {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						break
						//continue
					}
				}
				if isDelete {
					k := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					fmt.Printf("配置：%s被删除\n", k)
					t.taskMap[k].cancelFunc()
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

func GetNewConf() chan<- []*etcd.LogEntry {
	return taskMgr.newConfChan
}
