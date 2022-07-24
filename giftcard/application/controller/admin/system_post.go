package admin

import (
	"giftcard/application/controller"
	"giftcard/application/service"
	"giftcard/library/context/api"
	"giftcard/library/context/result"
)

//登录
type SystemPost struct {
	controller.BaseController
}

//管理员登录
func (this *SystemPost) PageList(ctx *api.Context) *result.Result {
	service := service.NewSystemPost()
	resp, err := service.PageList()
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}
