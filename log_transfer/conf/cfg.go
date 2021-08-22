package conf

type LogTransfer struct {
	Kafka `ini:"kafka"`
	ESCfg `ini:"es"`
}

type Kafka struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESCfg struct {
	Address string `ini:"address"`
}
