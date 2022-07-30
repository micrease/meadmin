package handler

import (
	"admease/app/system/dto"
	"admease/app/system/service"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
)

type SystemDictData struct {
	Handler
}

func (this *SystemDictData) List(ctx *api.Context) *result.Result {

	var req dto.SystemDictDataListReq
	validate.BindWithPanic(ctx, &req)

	service := service.NewSystemDictData()
	resp, err := service.List(ctx, req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}