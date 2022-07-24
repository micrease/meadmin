package models

import "time"

type SystemLoginLog struct {
	ID	uint64	`json:"id"`	// 主键
	Username	string	`json:"username"`	// 用户名
	Ip	string	`json:"ip"`	// 登录IP地址
	IpLocation	string	`json:"ip_location"`	// IP所属地
	Os	string	`json:"os"`	// 操作系统
	Browser	string	`json:"browser"`	// 浏览器
	Status	string	`json:"status"`	// 登录状态 (0成功 1失败)
	Message	string	`json:"message"`	// 提示消息
	LoginTime	time.Time	`json:"login_time"`	// 登录时间
	Remark	string	`json:"remark"`	// 备注
}

func (model SystemLoginLog) GetID() uint64 {
	return model.ID
}

func (model SystemLoginLog) TableName() string {
	return "system_login_log"
}
