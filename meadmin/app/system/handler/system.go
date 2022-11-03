package handler

import (
	"meadmin/app/system/dto"
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

// 系统管理
type System struct {
	Handler
}

// 管理员登录
func (this *System) Login(ctx *api.Context) *result.Result {
	var req dto.SystemLoginReq
	validate.BindWithPanic(ctx, &req)
	resp, respErr := service.NewSystemUser(ctx).Login(req)
	if respErr != nil {
		return result.ErrorResult(respErr)
	}
	return result.Success(resp)
}

func (this *System) Logout(ctx *api.Context) *result.Result {
	err := service.NewSystemUser(ctx).Logout()
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success()
}

// 获取管理员信息
func (this *System) GetInfo(ctx *api.Context) *result.Result {
	var req dto.SystemLoginReq
	validate.BindWithPanic(ctx, &req)

	resp, respErr := service.NewSystemUser(ctx).GetInfo()
	if respErr != nil {
		return result.ErrorResult(respErr)
	}
	return result.Success(resp)
}

// PageList
func (this *System) PageList(ctx *api.Context) *result.Result {
	resp, respErr := service.NewSystemUser(ctx).PageList(ctx)
	if respErr != nil {
		return result.ErrorResult(respErr)
	}
	return result.Success(resp)
}
