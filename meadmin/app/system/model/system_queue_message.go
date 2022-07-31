package model

import "time"

type SystemQueueMessage struct {
	ID          uint64    `json:"id"`           // 主键
	ContentType string    `json:"content_type"` // 内容类型
	Title       string    `json:"title"`        // 消息标题
	Content     string    `json:"content"`      // 消息内容
	SendBy      uint64    `json:"send_by"`      // 发送人
	CreatedBy   uint64    `json:"created_by"`   // 创建者
	UpdatedBy   uint64    `json:"updated_by"`   // 更新者
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
	Remark      string    `json:"remark"`       // 备注
}

func (model SystemQueueMessage) GetID() uint64 {
	return model.ID
}

func (model SystemQueueMessage) TableName() string {
	return "system_queue_message"
}
