package model

import "time"

type SystemDictData struct {
	ID        uint64    `json:"id"`         // 主键
	TypeId    uint64    `json:"type_id"`    // 字典类型ID
	Label     string    `json:"label"`      // 字典标签
	Value     string    `json:"value"`      // 字典值
	Code      string    `json:"code"`       // 字典标示
	Sort      int       `json:"sort"`       // 排序
	Status    string    `json:"status"`     // 状态 (0正常 1停用)
	CreatedBy uint64    `json:"created_by"` // 创建者
	UpdatedBy uint64    `json:"updated_by"` // 更新者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at"` // 删除时间
	Remark    string    `json:"remark"`     // 备注
}

func (model SystemDictData) GetID() uint64 {
	return model.ID
}

func (model SystemDictData) TableName() string {
	return "system_dict_data"
}
