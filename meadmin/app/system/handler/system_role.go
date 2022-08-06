package handler

import (
	"meadmin/app/system/dto"
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

type SystemRole struct {
	Handler
}

func (this SystemRole) Save(ctx *api.Context) *result.Result {
	var req dto.SystemRoleSaveReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemRole()
	err := service.Save(ctx, req)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success()
}

func (this SystemRole) Update(ctx *api.Context) *result.Result {
	var req dto.SystemRoleSaveReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemRole()
	err := service.Save(ctx, req)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success()
}

func (this SystemRole) Index(ctx *api.Context) *result.Result {
	resp, err := service.NewSystemRole().GetPageList(ctx)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(resp)
}

func (this SystemRole) List(ctx *api.Context) *result.Result {
	resp, err := service.NewSystemRole().GetList(ctx)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(resp)
}
