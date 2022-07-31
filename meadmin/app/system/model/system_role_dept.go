package model

type SystemRoleDept struct {
	ID     uint64 `json:"id"`      // 主键，角色ID
	RoleId uint64 `json:"role_id"` // 角色主键
	DeptId uint64 `json:"dept_id"` // 部门主键
}

func (model SystemRoleDept) GetID() uint64 {
	return model.ID
}

func (model SystemRoleDept) TableName() string {
	return "system_role_dept"
}
