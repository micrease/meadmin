package admin

import (
	"giftcard/application/controller"
	"giftcard/application/service"
	"giftcard/library/context/api"
	"giftcard/library/context/result"
)

//登录
type SystemMenu struct {
	controller.BaseController
}

//管理员登录
func (this *SystemMenu) Index(ctx *api.Context) *result.Result {
	service := service.NewSystemMenu()
	resp, err := service.GetMenuList(ctx)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}
