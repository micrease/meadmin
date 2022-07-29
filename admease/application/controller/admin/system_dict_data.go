package admin

import (
	"giftcard/application/controller"
	"giftcard/application/dto"
	"giftcard/application/service"
	"giftcard/library/context/api"
	"giftcard/library/context/result"
	"giftcard/library/validate"
)

type SystemDictData struct {
	controller.BaseController
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
