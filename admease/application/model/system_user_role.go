package models

type SystemUserRole struct {
	ID     uint64 `json:"id"`      // 用户ID，主键
	UserId uint64 `json:"user_id"` // 用户主键
	RoleId uint64 `json:"role_id"` // 角色主键
}

func (model SystemUserRole) GetID() uint64 {
	return model.ID
}

func (model SystemUserRole) TableName() string {
	return "system_user_role"
}
