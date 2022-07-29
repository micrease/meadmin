package admin

import (
	"admease/application/controller"
	"admease/application/dto"
	"admease/application/service"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
)

//系统管理
type System struct {
	controller.BaseController
}

//管理员登录
func (this *System) Login(ctx *api.Context) *result.Result {
	var req dto.SystemLoginReq
	validate.BindWithPanic(ctx, &req)
	resp, respErr := service.NewSystemUser().Login(req)
	if respErr != nil {
		return result.Failed(respErr)
	}
	return result.Success(resp)
}

func (this *System) Logout(ctx *api.Context) *result.Result {
	err := service.NewSystemUser().Logout(ctx)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

//获取管理员信息
func (this *System) GetInfo(ctx *api.Context) *result.Result {
	var req dto.SystemLoginReq
	validate.BindWithPanic(ctx, &req)

	resp, respErr := service.NewSystemUser().GetIInfo(ctx.JwtClaimData.UserId, ctx.JwtClaimData.IsSuperAdmin)
	if respErr != nil {
		return result.Failed(respErr)
	}
	return result.Success(resp)
}

//PageList
func (this *System) PageList(ctx *api.Context) *result.Result {
	var req dto.SystemLoginReq
	validate.BindWithPanic(ctx, &req)
	resp, respErr := service.NewSystemUser().PageList()
	if respErr != nil {
		return result.Failed(respErr)
	}
	return result.Success(resp)
}
