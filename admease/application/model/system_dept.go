package models

import "time"

type SystemDept struct {
	ID	uint64	`json:"id"`	// 主键
	ParentId	uint64	`json:"parent_id"`	// 父ID
	Level	string	`json:"level"`	// 组级集合
	Name	string	`json:"name"`	// 部门名称
	Leader	string	`json:"leader"`	// 负责人
	Phone	string	`json:"phone"`	// 联系电话
	Status	string	`json:"status"`	// 状态 (0正常 1停用)
	Sort	int	`json:"sort"`	// 排序
	CreatedBy	uint64	`json:"created_by"`	// 创建者
	UpdatedBy	uint64	`json:"updated_by"`	// 更新者
	CreatedAt	time.Time	`json:"created_at"`	// 创建时间
	UpdatedAt	time.Time	`json:"updated_at"`	// 更新时间
	DeletedAt	time.Time	`json:"deleted_at"`	// 删除时间
	Remark	string	`json:"remark"`	// 备注
}

func (model SystemDept) GetID() uint64 {
	return model.ID
}

func (model SystemDept) TableName() string {
	return "system_dept"
}
