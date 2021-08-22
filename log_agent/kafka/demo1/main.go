// kafka的连接，生产消费

package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {

	// tmp := []byte("fasdfasdf")
	// fmt.Println(tmp)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 1
	config.Metadata.Retry.Backoff = time.Second

	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log" // 发送的topic随便定
	msg.Value = sarama.StringEncoder("********\nthis is a web log***********\n")

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) //地址要对，不然会卡在这
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid: %v offset: %v\n", pid, offset) // 每发一条消息，offset加1

	pid, offset, err = client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid: %v offset: %v\n", pid, offset)
}
