package model

import "time"

// 代码生成业务字段信息表
type SettingGenerateColumns struct {
	ID            uint64    `json:"id"`             // 主键
	TableId       uint64    `json:"table_id"`       // 所属表ID
	ColumnName    string    `json:"column_name"`    // 字段名称
	ColumnComment string    `json:"column_comment"` // 字段注释
	ColumnType    string    `json:"column_type"`    // 字段类型
	IsPk          string    `json:"is_pk"`          // 0 非主键 1 主键
	IsRequired    string    `json:"is_required"`    // 0 非必填 1 必填
	IsInsert      string    `json:"is_insert"`      // 0 非插入字段 1 插入字段
	IsEdit        string    `json:"is_edit"`        // 0 非编辑字段 1 编辑字段
	IsList        string    `json:"is_list"`        // 0 非列表显示字段 1 列表显示字段
	IsQuery       string    `json:"is_query"`       // 0 非查询字段 1 查询字段
	QueryType     string    `json:"query_type"`     // 查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围
	ViewType      string    `json:"view_type"`      // 页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）
	DictType      string    `json:"dict_type"`      // 字典类型
	AllowRoles    string    `json:"allow_roles"`    // 允许查看该字段的角色
	Options       string    `json:"options"`        // 字段其他设置
	Sort          int       `json:"sort"`           // 排序
	CreatedBy     uint64    `json:"created_by"`     // 创建者
	UpdatedBy     uint64    `json:"updated_by"`     // 更新者
	CreatedAt     time.Time `json:"created_at"`     // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`     // 更新时间
	Remark        string    `json:"remark"`         // 备注
}

func (model SettingGenerateColumns) GetID() uint64 {
	return model.ID
}

func (model SettingGenerateColumns) TableName() string {
	return "setting_generate_columns"
}
