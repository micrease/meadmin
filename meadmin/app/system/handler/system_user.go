package handler

import (
	"admease/app/system/dto"
	"admease/app/system/service"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
)

//用户模块
type SystemUser struct {
	Handler
}

//用户列表
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
