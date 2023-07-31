package handler

import (
	"github.com/spf13/cast"
	"meadmin/app/system/service"
	"meadmin/app/system/vo"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
	"strings"
)

// 登录
type SystemUser struct {
	Handler
}

// ReadInfo 查看user数据
func (this *SystemUser) ReadInfo(ctx *api.Context) *result.Result {
	userId := ctx.GinCtx.Param("id")
	resp, respErr := service.NewSystemUser(ctx).ReadInfoById(ctx, cast.ToUint64(userId))
	if respErr != nil {
		return result.ErrorResult(respErr)
	}
	return result.Success(resp)
}

// ClearCache 清理缓存
func (this *SystemUser) ClearCache(ctx *api.Context) *result.Result {
	var req vo.SystemUserClearCacheReq
	validate.BindWithPanic(ctx, &req)
	// todo 清理缓存逻辑
	return result.Success()
}

// ChangeStatus 修改用户状态
func (this *SystemUser) ChangeStatus(ctx *api.Context) *result.Result {
	var req vo.SystemUserChangeStatusReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser(ctx).ChangeStatus(req.ID, req.Status)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

func (this *SystemUser) SaveUser(ctx *api.Context) *result.Result {
	var req vo.SystemUserSaveReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser(ctx).Save(ctx, req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

func (this *SystemUser) DeleteUser(ctx *api.Context) *result.Result {
	id := ctx.GinCtx.Param("id")
	ids := strings.Split(id, ",")
	err := service.NewSystemUser(ctx).Delete(ids)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

func (this *SystemUser) InitUserPassword(ctx *api.Context) *result.Result {
	id := ctx.GinCtx.Param("id")
	err := service.NewSystemUser(ctx).ResetPassword(cast.ToUint64(id))
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

func (this *SystemUser) SetHomePage(ctx *api.Context) *result.Result {
	var req vo.SystemUserSetHomePageReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser(ctx).SetHomePage(req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

func (this *SystemUser) Update(ctx *api.Context) *result.Result {
	var req vo.SystemUserSaveReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewSystemUser(ctx).Save(ctx, req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

func (this *SystemUser) setHomePage(ctx *api.Context) *result.Result {
	return result.Success()
}
