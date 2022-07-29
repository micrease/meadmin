package handler

import (
	"admease/app/system/dto"
	"admease/app/system/service"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
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
