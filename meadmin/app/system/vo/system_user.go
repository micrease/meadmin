package vo

import "meadmin/app/system/model"

type SystemUserInfoResp struct {
	User
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
	PageQuery
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

type SystemUserSaveReq struct {
	ID       uint64   `form:"id" json:"id"`
	Avatar   string   `form:"avatar" json:"avatar"`
	DeptID   uint64   `form:"dept_id" json:"dept_id"`
	Email    string   `form:"email" json:"email"`
	NickName string   `form:"nickname" json:"nickname"`
	UserName string   `form:"username" json:"username"`
	Password string   `form:"password" json:"password"`
	Phone    string   `form:"phone" json:"phone"`
	PostIds  []uint64 `form:"post_ids" json:"post_ids"`
	Remark   string   `form:"remark" json:"remark"`
	RoleIds  []uint64 `form:"role_ids" json:"role_ids"`
	Status   string   `form:"status" json:"status"`
}

type SystemUserSetHomePageReq struct {
	ID        uint64 `form:"id" json:"id"`
	Dashboard string `form:"dashboard" json:"dashboard"`
}
