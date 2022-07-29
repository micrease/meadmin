package model

import "time"

type SystemNotice struct {
	ID        uint64    `json:"id"`         // 主键
	MessageId uint64    `json:"message_id"` // 消息ID
	Title     string    `json:"title"`      // 标题
	Type      string    `json:"type"`       // 公告类型（1通知 2公告）
	Content   string    `json:"content"`    // 公告内容
	ClickNum  uint      `json:"click_num"`  // 浏览次数
	CreatedBy uint64    `json:"created_by"` // 创建者
	UpdatedBy uint64    `json:"updated_by"` // 更新者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at"` // 删除时间
	Remark    string    `json:"remark"`     // 备注
}

func (model SystemNotice) GetID() uint64 {
	return model.ID
}

func (model SystemNotice) TableName() string {
	return "system_notice"
}
