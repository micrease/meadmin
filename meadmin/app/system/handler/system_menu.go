package handler

import (
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
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
