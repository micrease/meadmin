package consumer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"log"
	"meadmin/library/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/**
KAFKA_BROKER_LIST=172.20.1.248:9092
KAFKA_BROKER_VERSION=1.0.0
//存款topic固定
KAFKA_DEPOSIT_TOPIC=deposit
//组id 自定义
KAFKA_DEPOSIT_GROUP_ID=laborday
//投注topic固定
KAFKA_BET_TOPIC=order
//组id 自定义
KAFKA_BET_GROUP_ID=laborday
*/

func ShowMetadata(addrs []string) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		return
	}
	defer client.Close()
	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))
	for _, topic := range topics {
		logger.Info("ShowMetadata", zap.String("topic", topic))
	}

	// get broker set
	brokers := client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("%s\n", broker.Addr())
	}
}

func MessageHandle(message *sarama.ConsumerMessage) {
	msg := fmt.Sprintf("Message claimed: value = %s, timestamp = %v, topic = %s, partions = %d, offset = %d", string(message.Value), message.Timestamp, message.Topic, message.Partition, message.Offset)
	logger.Info(msg)
}

type KafkaConsumerGroupAction struct {
	group  sarama.ConsumerGroup
	handle func()
}

func NewKafkaConsumerGroupAction(brokers []string, groupId string) *KafkaConsumerGroupAction {
	logger.Info("NewKafkaConsumerGroupAction", zap.String("groupId", groupId), zap.Any("brokers", brokers))
	ShowMetadata(brokers)
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stdout, "[consumer_group]", log.Lshortfile)
	// 重平衡策略
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second
	config.Consumer.IsolationLevel = sarama.ReadCommitted
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Version = sarama.V2_7_0_0
	consumerGroup, e := sarama.NewConsumerGroup(brokers, groupId, config)
	if e != nil {
		logger.Info("NewKafkaConsumerGroupAction", zap.Error(e))
		return nil
	}
	return &KafkaConsumerGroupAction{group: consumerGroup}
}

func (K *KafkaConsumerGroupAction) Consume(topics []string, handle_func func(message *sarama.ConsumerMessage)) {

	if K == nil {
		return
	}
	var wg sync.WaitGroup
	ctx := context.Background()
	var consumer = KafkaConsumerGroupHandler{ready: make(chan bool), handle: handle_func}
	go func() {
		defer wg.Done()
		for {
			if err := K.group.Consume(ctx, topics, &consumer); err != nil {
				logger.Info("consumer|Error from consumer", zap.Error(err))
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
	logger.Info("Sarama consumer up and running!...")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		logger.Info("consumer|terminating: context cancelled")
	case <-sigterm:
		logger.Info("consumer|terminating: via signal")
	}
	wg.Wait()
	if err := K.group.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

type KafkaConsumerGroupHandler struct {
	ready  chan bool
	handle func(message *sarama.ConsumerMessage)
}

func (K *KafkaConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (K *KafkaConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
func (K *KafkaConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		K.handle(message)
		lag := claim.HighWaterMarkOffset() - message.Offset
		fmt.Println(lag)
		session.MarkMessage(message, "")
	}

	return nil
}
