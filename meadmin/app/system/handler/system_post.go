package handler

import (
	"meadmin/app/system/dto"
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

type SystemPost struct {
	Handler
}

func (this *SystemPost) PageList(ctx *api.Context) *result.Result {
	service := service.NewSystemPost()
	resp, err := service.PageList()
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}

func (this *SystemPost) PostList(ctx *api.Context) *result.Result {
	service := service.NewSystemPost()
	resp, err := service.PostList()
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}

func (this *SystemPost) Save(ctx *api.Context) *result.Result {
	var req dto.SystemPostSaveReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemPost()
	err := service.Save(ctx, req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

func (this SystemPost) Update(ctx *api.Context) *result.Result {
	var req dto.SystemPostSaveReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemPost()
	err := service.Save(ctx, req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}
