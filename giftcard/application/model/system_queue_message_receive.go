package models

type SystemQueueMessageReceive struct {
	ID         uint64 `json:"id"`          // 主键
	MessageId  uint64 `json:"message_id"`  // 队列消息主键
	UserId     uint64 `json:"user_id"`     // 接收用户主键
	ReadStatus string `json:"read_status"` // 已读状态 (0未读 1已读)
}

func (model SystemQueueMessageReceive) GetID() uint64 {
	return model.ID
}

func (model SystemQueueMessageReceive) TableName() string {
	return "system_queue_message_receive"
}
