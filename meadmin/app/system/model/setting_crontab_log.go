package model

import "time"

// 定时任务执行日志表
type SettingCrontabLog struct {
	ID            uint64    `json:"id"`             // 主键
	CrontabId     uint64    `json:"crontab_id"`     // 任务ID
	Name          string    `json:"name"`           // 任务名称
	Target        string    `json:"target"`         // 任务调用目标字符串
	Parameter     string    `json:"parameter"`      // 任务调用参数
	ExceptionInfo string    `json:"exception_info"` // 异常信息
	Status        string    `json:"status"`         // 执行状态 (0成功 1失败)
	CreatedAt     time.Time `json:"created_at"`     // 创建时间
}

func (model SettingCrontabLog) GetID() uint64 {
	return model.ID
}

func (model SettingCrontabLog) TableName() string {
	return "setting_crontab_log"
}
