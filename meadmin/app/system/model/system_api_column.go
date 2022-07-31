package model

import "time"

type SystemApiColumn struct {
	ID           uint64    `json:"id"`            // 主键
	ApiId        uint64    `json:"api_id"`        // 接口主键
	Name         string    `json:"name"`          // 字段名称
	Type         string    `json:"type"`          // 字段类型 0 请求 1 返回
	DataType     string    `json:"data_type"`     // 数据类型
	IsRequired   string    `json:"is_required"`   // 是否必填 0 非必填 1 必填
	DefaultValue string    `json:"default_value"` // 默认值
	Status       string    `json:"status"`        // 状态 (0正常 1停用)
	Description  string    `json:"description"`   // 字段说明
	CreatedBy    uint64    `json:"created_by"`    // 创建者
	UpdatedBy    uint64    `json:"updated_by"`    // 更新者
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
	DeletedAt    time.Time `json:"deleted_at"`    // 删除时间
	Remark       string    `json:"remark"`        // 备注
}

func (model SystemApiColumn) GetID() uint64 {
	return model.ID
}

func (model SystemApiColumn) TableName() string {
	return "system_api_column"
}
