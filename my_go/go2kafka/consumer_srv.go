package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func ConsumerSrv(hostList []string, topic string) {
	fmt.Println("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer(hostList, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s, source_value:%v\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value), msg)
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}

func main() {
	fmt.Println("vim-go")
	hostList := []string{
		"localhost:9092",
	}
	topic := "test"

	ConsumerSrv(hostList, topic)
}
