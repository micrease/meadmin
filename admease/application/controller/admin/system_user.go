package admin

import (
	"admease/application/controller"
	"admease/application/dto"
	"admease/application/service"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
)

//登录
type SystemUser struct {
	controller.BaseController
}

//管理员登录
func (this *SystemUser) Index(ctx *api.Context) *result.Result {
	var req dto.SystemLoginReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemUser()
	resp, em := service.Login(req)
	if em != nil {
		return result.Failed(em)
	}
	return result.Success(resp)
}
