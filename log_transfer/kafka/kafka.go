package kafka

import (
	"fmt"
	"log_transfer/es"

	"github.com/Shopify/sarama"
)

var pc sarama.PartitionConsumer

// 从kafka消费数据
func Init(address []string, topic string) error {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}

	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println("分区列表：", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		fmt.Println("进入分区：...")
		// 针对每个分区创建一个对应的分区消费者
		pc, err = consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		// defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			fmt.Println("enter goroutine...")
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				es.SendToEs(topic, msg.Value)
			}
		}(pc)
	}
	// for {
	// }
	return nil
}
