package handler

import (
	"admease/app/system/service"
	"admease/library/context/api"
	"admease/library/context/result"
)

type SystemMenu struct {
	Handler
}

func (this *SystemMenu) Index(ctx *api.Context) *result.Result {
	service := service.NewSystemMenu()
	resp, err := service.GetMenuList(ctx)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}
