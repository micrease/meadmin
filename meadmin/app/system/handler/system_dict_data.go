package handler

import (
	"meadmin/app/system/service"
	"meadmin/app/system/vo"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

type SystemDictData struct {
	Handler
}

func (this *SystemDictData) List(ctx *api.Context) *result.Result {

	var req vo.SystemDictDataListReq
	validate.BindWithPanic(ctx, &req)

	service := service.NewSystemDictData()
	resp, err := service.List(ctx, req)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(resp)
}
