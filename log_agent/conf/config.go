package conf

type AppConf struct {
	KafkaConf   `ini:"kafka"`
	EtcdConf    `ini:"etcd"`
	TaillogConf `ini:"tail"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	Topic       string `ini:"topic"`
	ChanMaxSize int    `ini:"chan_maxsize"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"key"` // etcd管理配置的key， value是json化的配置项
	Timeout int    `ini:"timeout"`
}

type TaillogConf struct {
	FileName string `ini:"filename"`
}
