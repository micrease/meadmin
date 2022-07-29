package admin

import (
	"admease/application/controller"
	"admease/application/service"
	"admease/library/context/api"
	"admease/library/context/result"
)

type SystemMenu struct {
	controller.BaseController
}

func (this *SystemMenu) Index(ctx *api.Context) *result.Result {
	service := service.NewSystemMenu()
	resp, err := service.GetMenuList(ctx)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}
