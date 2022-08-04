package handler

import (
	"admease/app/system/dto"
	"admease/app/system/service"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
	"github.com/spf13/cast"
	"strings"
)

//登录
type SystemUser struct {
	Handler
}

//管理员登录
func (u *SystemUser) Index(ctx *api.Context) *result.Result {
	var req dto.SystemLoginReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemUser()
	resp, em := service.Login(req)
	if em != nil {
		return result.Failed(em)
	}
	return result.Success(resp)
}

// ReadInfo 查看user数据
func (u *SystemUser) ReadInfo(ctx *api.Context) *result.Result {
	userId := ctx.GinCtx.Param("id")
	resp, respErr := service.NewSystemUser().ReadInfoById(ctx, cast.ToUint64(userId))

	if respErr != nil {
		return result.Failed(respErr)
	}
	return result.Success(resp)
}

// ClearCache 清理缓存
func (u *SystemUser) ClearCache(ctx *api.Context) *result.Result {
	var req dto.SystemUserClearCacheReq
	validate.BindWithPanic(ctx, &req)
	// todo 清理缓存逻辑
	return result.Success()
}

// ChangeStatus 修改用户状态
func (u *SystemUser) ChangeStatus(ctx *api.Context) *result.Result {
	var req dto.SystemUserChangeStatusReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser().ChangeStatus(req.ID, req.Status)
	if err != nil {
		return result.Failed(err)
	}
	return result.Success()
}

func (u *SystemUser) SaveUser(ctx *api.Context) *result.Result {
	var req dto.SystemUserSaveReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser().Save(ctx, req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

func (u SystemUser) DeleteUser(ctx *api.Context) *result.Result {
	id := ctx.GinCtx.Param("id")
	ids := strings.Split(id, ",")
	err := service.NewSystemUser().Delete(ids)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

func (u SystemUser) InitUserPassword(ctx *api.Context) *result.Result {
	id := ctx.GinCtx.Param("id")
	err := service.NewSystemUser().ResetPassword(cast.ToUint64(id))
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

func (u SystemUser) SetHomePage(ctx *api.Context) *result.Result {
	var req dto.SystemUserSetHomePageReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser().SetHomePage(req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

func (u *SystemUser) Update(ctx *api.Context) *result.Result {
	var req dto.SystemUserSaveReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser().Save(ctx, req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}