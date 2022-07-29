package dto

import "admease/library/logger"

//所有的请求类对象，都要继承一下这个
type ReqHeader struct {
	TraceId string     `json:"-"`
	Log     logger.Log `json:"-"`
}

//分页类对象继承
type PageQuery struct {
	PageSize int `json:"page_size" form:"page_size" ` //分页:每页条数
	PageNo   int `json:"page_no"   form:"page_no" `   //分页:当前页码,起始页为1
}

//分页返回数据类型
type PageListResp[T any] struct {
	PageSize  int   `json:"page_size"`  //每页条数
	PageNo    int   `json:"page_no"`    //当前页码
	TotalPage int64 `json:"total_page"` //总页数
	TotalSize int64 `json:"total_size"` //总条数
	List      []T   `json:"list"`       //数据列表
}

//system分页
type PageInfo struct {
	CurrentPage int   `json:"currentPage"`
	Total       int64 `json:"total"`
	TotalPage   int64 `json:"totalPage"`
}

//分页返回数据类型
type SystemPageListResp[T any] struct {
	PageInfo PageInfo `json:"pageInfo"`
	Items    []*T     `json:"items"` //数据列表
}
