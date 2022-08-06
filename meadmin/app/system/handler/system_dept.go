package handler

import (
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
)

//登录
type SystemDept struct {
	Handler
}

//管理员登录
func (this *SystemDept) GetListTree(ctx *api.Context) *result.Result {
	service := service.NewSystemDept()
	resp, err := service.GetListTree()
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(resp)
}

//Index
func (this *SystemDept) Index(ctx *api.Context) *result.Result {
	service := service.NewSystemDept()
	resp, err := service.GetModelListTree()
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(resp)
}
