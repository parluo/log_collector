package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

var (
	gClient     sarama.SyncProducer
	logDataChan chan *logData
)

func KafakInstance() *sarama.SyncProducer {
	return &gClient
}
func Init(address string, kafkaMaxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	gClient, err = sarama.NewSyncProducer([]string{address}, config)
	if err != nil {
		fmt.Println("kafka init: fail to init. producer closed, err:", err)
		return err
	}
	logDataChan = make(chan *logData, kafkaMaxSize)
	go sendToKafka()
	return nil
}
func sendToKafka() {
	// msgData := &logData{}
	for {
		select {
		case msgData := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = msgData.topic // 发送的topic随便定
			msg.Value = sarama.StringEncoder(msgData.data)

			pit, offset, err := gClient.SendMessage(msg)
			if err != nil {
				fmt.Println("send message to kafka failed!")
				continue
			}
			fmt.Printf("successful send msg to kafka: pit:%v, offset:%v\n", pit, offset)
		default:
			time.Sleep(time.Second)
		}
	}
}

func SendToChan(topic, data string) {
	logDataChan <- &logData{topic, data}
}
