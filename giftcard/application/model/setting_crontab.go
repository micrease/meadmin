package models

import "time"

// 定时任务信息表
type SettingCrontab struct {
	ID	uint64	`json:"id"`	// 主键
	Name	string	`json:"name"`	// 任务名称
	Type	string	`json:"type"`	// 任务类型 (1 command, 2 class, 3 url, 4 eval)
	Target	string	`json:"target"`	// 调用任务字符串
	Parameter	string	`json:"parameter"`	// 调用任务参数
	Rule	string	`json:"rule"`	// 任务执行表达式
	Singleton	string	`json:"singleton"`	// 是否单次执行 (0 是 1 不是)
	Status	string	`json:"status"`	// 状态 (0正常 1停用)
	CreatedBy	uint64	`json:"created_by"`	// 创建者
	UpdatedBy	uint64	`json:"updated_by"`	// 更新者
	CreatedAt	time.Time	`json:"created_at"`	// 创建时间
	UpdatedAt	time.Time	`json:"updated_at"`	// 更新时间
	Remark	string	`json:"remark"`	// 备注
}

func (model SettingCrontab) GetID() uint64 {
	return model.ID
}

func (model SettingCrontab) TableName() string {
	return "setting_crontab"
}
