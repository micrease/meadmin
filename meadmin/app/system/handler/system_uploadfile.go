package handler

import (
	"meadmin/app/system/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
)

type SystemUploadfile struct {
	Handler
}

func (this *SystemUploadfile) Upload(ctx *api.Context) *result.Result {
	//获取文件头
	file, err := ctx.GinCtx.FormFile("image")
	if err != nil {
		return result.ErrorMessage(err)
	}
	model, err := service.NewSystemUploadFile().Uploadfile(ctx, file)
	//获取文件名
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(model)
}

func (this *SystemUploadfile) Index(ctx *api.Context) *result.Result {
	resp, err := service.NewSystemUploadFile().GetPageList(ctx)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(resp)
}
