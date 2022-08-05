package handler

import (
	"meadmin/app/system/dto"
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
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
