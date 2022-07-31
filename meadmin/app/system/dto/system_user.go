package dto

import "admease/app/system/model"

type SystemUserInfoResp struct {
	UserDTO
	RoleList []model.SystemRole `json:"roleList"`
	PostList []int              `json:"postList"`
}
