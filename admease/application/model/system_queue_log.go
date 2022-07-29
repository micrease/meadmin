package models

import "time"

type SystemQueueLog struct {
	ID	uint64	`json:"id"`	// 主键
	ExchangeName	string	`json:"exchange_name"`	// 交换机名称
	RoutingKeyName	string	`json:"routing_key_name"`	// 路由名称
	QueueName	string	`json:"queue_name"`	// 队列名称
	QueueContent	string	`json:"queue_content"`	// 队列数据
	LogContent	string	`json:"log_content"`	// 队列日志
	ProduceStatus	string	`json:"produce_status"`	// 生产状态 0:未生产 1:生产中 2:生产成功 3:生产失败 4:生产重复
	ConsumeStatus	string	`json:"consume_status"`	// 消费状态 0:未消费 1:消费中 2:消费成功 3:消费失败 4:消费重复
	DelayTime	uint	`json:"delay_time"`	// 延迟时间（秒）
	CreatedBy	uint64	`json:"created_by"`	// 创建者
	UpdatedBy	uint64	`json:"updated_by"`	// 更新者
	CreatedAt	time.Time	`json:"created_at"`	// 创建时间
	UpdatedAt	time.Time	`json:"updated_at"`	// 更新时间
}

func (model SystemQueueLog) GetID() uint64 {
	return model.ID
}

func (model SystemQueueLog) TableName() string {
	return "system_queue_log"
}
