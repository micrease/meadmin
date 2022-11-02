package model

import "time"

type SystemRole struct {
	ID        uint64    `json:"id"`         // 主键，角色ID
	Name      string    `json:"name"`       // 角色名称
	Code      string    `json:"code"`       // 角色代码
	DataScope string    `json:"data_scope"` // 数据范围（0：全部数据权限 1：自定义数据权限 2：本部门数据权限 3：本部门及以下数据权限 4：本人数据权限）
	Status    int       `json:"status"`     // 状态 (0正常 1停用)
	Sort      int       `json:"sort"`       // 排序
	CreatedBy uint64    `json:"created_by"` // 创建者
	UpdatedBy uint64    `json:"updated_by"` // 更新者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	Remark    string    `json:"remark"`     // 备注
}

func (model SystemRole) GetID() any {
	return model.ID
}

func (model SystemRole) TableName() string {
	return "system_role"
}
