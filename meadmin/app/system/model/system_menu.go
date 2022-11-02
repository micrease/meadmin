package model

import "time"

const (
	StatusEnable  = 0
	StatusDisable = 1
)

type SystemMenu struct {
	ID        uint64    `json:"id"`         // 主键
	ParentId  uint64    `json:"parent_id"`  // 父ID
	Level     string    `json:"level"`      // 组级集合
	Name      string    `json:"name"`       // 菜单名称
	Code      string    `json:"code"`       // 菜单标识代码
	Icon      string    `json:"icon"`       // 菜单图标
	Route     string    `json:"route"`      // 路由地址
	Component string    `json:"component"`  // 组件路径
	Redirect  string    `json:"redirect"`   // 跳转地址
	IsHidden  int       `json:"is_hidden"`  // 是否隐藏 (0是 1否)
	Type      string    `json:"type"`       // 菜单类型, (M菜单 B按钮 L链接 I iframe)
	Status    string    `json:"status"`     // 状态 (0正常 1停用)
	Sort      int       `json:"sort"`       // 排序
	CreatedBy uint64    `json:"created_by"` // 创建者
	UpdatedBy uint64    `json:"updated_by"` // 更新者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at"` // 删除时间
	Remark    string    `json:"remark"`     // 备注
}

func (model SystemMenu) GetID() any {
	return model.ID
}

func (model SystemMenu) TableName() string {
	return "system_menu"
}
