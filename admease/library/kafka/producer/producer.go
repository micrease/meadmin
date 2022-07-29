package producer

import (
	"admease/library/logger"
	"admease/system/config"
	"fmt"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"log"
	"strings"
	"time"
)

var producer sarama.SyncProducer

func GetProducer() sarama.SyncProducer {
	if producer != nil {
		return producer
	}
	conf := config.GetConfig()
	//kafka地址
	addrArr := strings.Split(conf.Kafka.YuanbaoAddrs, ",")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	var err error
	producer, err = sarama.NewSyncProducer(addrArr, config)
	if err != nil {
		logger.Error("Kafka GetProducer Error", zap.Error(err))
	}
	return producer
}

func PushMessage(topic string, data string) {
	producer := GetProducer()
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 连接 kafka
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		logger.Error("Kafka PushMessage Error", zap.Error(err))
		return
	}
	logger.Info("Kafka PushMessage", zap.Int32("message", partition), zap.Int64("offset", offset))
}

func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	msg := &sarama.ProducerMessage{}
	msg.Topic = "example"
	msg.Value = sarama.StringEncoder("xxxxx23423423423423423")

	// 连接 kafka
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Print(err)
		return
	}
	defer producer.Close()
	for {

		// 发送消息
		message, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(message, "====", offset)
		time.Sleep(time.Second * 2)
	}
}
