package dto

import "admease/app/system/model"

type SystemUserInfoResp struct {
	UserDTO
	RoleList []model.SystemRole `json:"roleList"`
	PostList []int              `json:"postList"`
}

type SystemUserClearCacheReq struct {
	ID uint64 `json:"id"`
}

type SystemUserChangeStatusReq struct {
	ID     uint64 `json:"id"`
	Status string `json:"status"`
}

type SystemUserListReq struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
	OrderBy   string `form:"orderBy"`
	OrderType string `form:"orderType"`
	UserName  string `form:"username"`
	NickName  string `form:"nickname"`
	Phone     string `form:"phone"`
	Email     string `form:"email"`
	MaxDate   string `form:"maxDate"`
	MinDate   string `form:"minDate"`
	Status    string `form:"status"`
}
