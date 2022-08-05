package dto

import "meadmin/app/system/model"

type SystemUserInfoResp struct {
	User
	RoleList []model.SystemRole `json:"roleList"`
	PostList []int              `json:"postList"`
}
