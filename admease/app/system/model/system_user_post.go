package model

type SystemUserPost struct {
	ID     uint64 `json:"id"`      // 用户ID，主键
	UserId uint64 `json:"user_id"` // 用户主键
	PostId uint64 `json:"post_id"` // 岗位主键
}

func (model SystemUserPost) GetID() uint64 {
	return model.ID
}

func (model SystemUserPost) TableName() string {
	return "system_user_post"
}
