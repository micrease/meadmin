package models

import "time"

// 代码生成业务信息表
type SettingGenerateTables struct {
	ID            uint64    `json:"id"`             // 主键
	TableNamee    string    `json:"table_name"`     // 表名称
	TableComment  string    `json:"table_comment"`  // 表注释
	ModuleName    string    `json:"module_name"`    // 所属模块
	Namespace     string    `json:"namespace"`      // 命名空间
	MenuName      string    `json:"menu_name"`      // 生成菜单名
	BelongMenuId  uint64    `json:"belong_menu_id"` // 所属菜单
	PackageName   string    `json:"package_name"`   // 控制器包名
	Type          string    `json:"type"`           // 生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD
	GenerateType  string    `json:"generate_type"`  // 0 压缩包下载 1 生成到模块
	GenerateMenus string    `json:"generate_menus"` // 生成菜单列表
	BuildMenu     string    `json:"build_menu"`     // 是否构建菜单
	ComponentType string    `json:"component_type"` // 组件显示方式
	Options       string    `json:"options"`        // 其他业务选项
	CreatedBy     uint64    `json:"created_by"`     // 创建者
	UpdatedBy     uint64    `json:"updated_by"`     // 更新者
	CreatedAt     time.Time `json:"created_at"`     // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`     // 更新时间
	Remark        string    `json:"remark"`         // 备注
}

func (model SettingGenerateTables) GetID() uint64 {
	return model.ID
}

func (model SettingGenerateTables) TableName() string {
	return "setting_generate_tables"
}
