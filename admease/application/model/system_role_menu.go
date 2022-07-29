package models

type SystemRoleMenu struct {
	ID     uint64 `json:"id"`      // 主键，角色ID
	RoleId uint64 `json:"role_id"` // 角色主键
	MenuId uint64 `json:"menu_id"` // 菜单主键
}

func (model SystemRoleMenu) GetID() uint64 {
	return model.ID
}

func (model SystemRoleMenu) TableName() string {
	return "system_role_menu"
}
