package model

import "time"

type SystemOperLog struct {
	ID           uint64    `json:"id"`            // 主键
	Username     string    `json:"username"`      // 用户名
	Method       string    `json:"method"`        // 请求方式
	Router       string    `json:"router"`        // 请求路由
	ServiceName  string    `json:"service_name"`  // 业务名称
	Ip           string    `json:"ip"`            // 请求IP地址
	IpLocation   string    `json:"ip_location"`   // IP所属地
	RequestData  string    `json:"request_data"`  // 请求数据
	ResponseCode string    `json:"response_code"` // 响应状态码
	ResponseData string    `json:"response_data"` // 响应数据
	CreatedBy    uint64    `json:"created_by"`    // 创建者
	UpdatedBy    uint64    `json:"updated_by"`    // 更新者
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
	DeletedAt    time.Time `json:"deleted_at"`    // 删除时间
	Remark       string    `json:"remark"`        // 备注
}

func (model SystemOperLog) GetID() uint64 {
	return model.ID
}

func (model SystemOperLog) TableName() string {
	return "system_oper_log"
}
