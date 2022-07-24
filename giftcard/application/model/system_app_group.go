package models

import "time"

type SystemAppGroup struct {
	ID	uint64	`json:"id"`	// 主键
	Name	string	`json:"name"`	// 应用组名称
	Status	string	`json:"status"`	// 状态 (0正常 1停用)
	CreatedBy	uint64	`json:"created_by"`	// 创建者
	UpdatedBy	uint64	`json:"updated_by"`	// 更新者
	CreatedAt	time.Time	`json:"created_at"`	// 创建时间
	UpdatedAt	time.Time	`json:"updated_at"`	// 更新时间
	DeletedAt	time.Time	`json:"deleted_at"`	// 删除时间
	Remark	string	`json:"remark"`	// 备注
}

func (model SystemAppGroup) GetID() uint64 {
	return model.ID
}

func (model SystemAppGroup) TableName() string {
	return "system_app_group"
}
