package model

import "time"

type SystemApp struct {
	ID          uint64    `json:"id"`          // 主键
	GroupId     uint64    `json:"group_id"`    // 应用组ID
	AppName     string    `json:"app_name"`    // 应用名称
	AppId       string    `json:"app_id"`      // 应用ID
	AppSecret   string    `json:"app_secret"`  // 应用密钥
	Status      string    `json:"status"`      // 状态 (0正常 1停用)
	Description string    `json:"description"` // 应用介绍
	CreatedBy   uint64    `json:"created_by"`  // 创建者
	UpdatedBy   uint64    `json:"updated_by"`  // 更新者
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`  // 更新时间
	DeletedAt   time.Time `json:"deleted_at"`  // 删除时间
	Remark      string    `json:"remark"`      // 备注
}

func (model SystemApp) GetID() uint64 {
	return model.ID
}

func (model SystemApp) TableName() string {
	return "system_app"
}
