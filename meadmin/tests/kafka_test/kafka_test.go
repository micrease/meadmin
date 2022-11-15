package kafka_test

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang-module/carbon/v2"
	"meadmin/library/kafka/consumer"
	"meadmin/library/kafka/producer"
	"meadmin/library/logger"
	"meadmin/system/config"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	go ConsumerMessage()
	time.Sleep(time.Second * 2)
	for i := 0; i < 10000; i++ {
		now := carbon.Now().ToDateTimeString()
		msg := fmt.Sprintf("发送消息_%s_%d", now, i)
		producer.PushMessage("test", msg)
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func MessageHandle(message *sarama.ConsumerMessage) {
	msg := fmt.Sprintf("Message claimed: value = %s, timestamp = %v, topic = %s, partions = %d, offset = %d", string(message.Value), message.Timestamp, message.Topic, message.Partition, message.Offset)
	fmt.Println(msg)
}

func ConsumerMessage() {
	conf := config.GetConfig()
	group := consumer.NewKafkaConsumerGroupAction([]string{conf.KafkaBroker}, "group_1")
	group.Consume([]string{"test"}, MessageHandle)
}

func TestMain(m *testing.M) {
	fmt.Println("run main")
	//解析配置文件
	config.InitConfig("../../resources/config-dev.yaml")
	//Log初始化
	logger.InitLog()
	m.Run()
}
