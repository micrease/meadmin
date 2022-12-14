package model

import "time"

type SystemPost struct {
	ID        uint64    `json:"id"`         // 主键
	Name      string    `json:"name"`       // 岗位名称
	Code      string    `json:"code"`       // 岗位代码
	Sort      int       `json:"sort"`       // 排序
	Status    string    `json:"status"`     // 状态 (0正常 1停用)
	CreatedBy uint64    `json:"created_by"` // 创建者
	UpdatedBy uint64    `json:"updated_by"` // 更新者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	Remark    string    `json:"remark"`     // 备注
}

func (model SystemPost) GetID() any {
	return model.ID
}

func (model SystemPost) TableName() string {
	return "system_post"
}
