package admin

import (
	"giftcard/application/controller"
	"giftcard/application/service"
	"giftcard/library/context/api"
	"giftcard/library/context/result"
)

type SystemUploadfile struct {
	controller.BaseController
}

func (this *SystemUploadfile) Upload(ctx *api.Context) *result.Result {
	//获取文件头
	file, err := ctx.GinCtx.FormFile("image")
	if err != nil {
		return result.ServerError(err)
	}
	model, err := service.NewSystemUploadFile().Uploadfile(ctx, file)
	//获取文件名
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(model)
}

func (this *SystemUploadfile) Index(ctx *api.Context) *result.Result {
	resp, err := service.NewSystemUploadFile().GetPageList(ctx)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}
