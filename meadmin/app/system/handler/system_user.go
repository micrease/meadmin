package handler

import (
	"github.com/spf13/cast"
	"meadmin/app/system/dto"
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

//登录
type SystemUser struct {
	Handler
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

func (this *System) ReadInfo(ctx *api.Context) *result.Result {
	userId := ctx.GinCtx.Param("id")
	resp, respErr := service.NewSystemUser().ReadInfoById(ctx, cast.ToUint64(userId))

	if respErr != nil {
		return result.Failed(respErr)
	}
	return result.Success(resp)
}
