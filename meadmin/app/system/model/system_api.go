package model

import "time"

type SystemApi struct {
	ID          uint64    `json:"id"`           // 主键
	GroupId     uint64    `json:"group_id"`     // 接口组ID
	Name        string    `json:"name"`         // 接口名称
	AccessName  string    `json:"access_name"`  // 接口访问名称
	ClassName   string    `json:"class_name"`   // 类命名空间
	MethodName  string    `json:"method_name"`  // 方法名
	AuthMode    string    `json:"auth_mode"`    // 认证模式 (0简易 1复杂)
	RequestMode string    `json:"request_mode"` // 请求模式 (A 所有 P POST G GET)
	Description string    `json:"description"`  // 接口说明介绍
	Response    string    `json:"response"`     // 返回内容示例
	Status      string    `json:"status"`       // 状态 (0正常 1停用)
	CreatedBy   uint64    `json:"created_by"`   // 创建者
	UpdatedBy   uint64    `json:"updated_by"`   // 更新者
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
	DeletedAt   time.Time `json:"deleted_at"`   // 删除时间
	Remark      string    `json:"remark"`       // 备注
}

func (model SystemApi) GetID() uint64 {
	return model.ID
}

func (model SystemApi) TableName() string {
	return "system_api"
}
