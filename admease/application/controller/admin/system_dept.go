package admin

import (
	"giftcard/application/controller"
	"giftcard/application/service"
	"giftcard/library/context/api"
	"giftcard/library/context/result"
)

//登录
type SystemDept struct {
	controller.BaseController
}

//管理员登录
func (this *SystemDept) GetListTree(ctx *api.Context) *result.Result {
	service := service.NewSystemDept()
	resp, err := service.GetListTree()
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}

//Index
func (this *SystemDept) Index(ctx *api.Context) *result.Result {
	service := service.NewSystemDept()
	resp, err := service.GetModelListTree()
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}
