package dto

type SystemRoleSaveReq struct {
	ID     uint64 `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Code   string `form:"code" json:"code"`
	Remark string `form:"remark" json:"remark"`
	Sort   int    `form:"sort" json:"sort"`
	Status string `form:"status" json:"status"`
}

type SystemRoleListReq struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
	OrderBy   string `form:"orderBy"`
	OrderType string `form:"orderType"`
	Name      string `form:"name"`
	Code      string `form:"code"`
	MaxDate   string `form:"maxDate"`
	MinDate   string `form:"minDate"`
	Status    string `form:"status"`
}
