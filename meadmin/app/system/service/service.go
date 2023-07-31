package service

import (
	"github.com/micrease/gorme"
	"go.uber.org/zap"
	"meadmin/app/system/vo"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/logger"
)

type Service struct {
	ctx     *api.Context
	TraceId string
	Log     logger.Log
}

func (service *Service) Init(ctx *api.Context) *Service {
	service.TraceId = ctx.TraceId
	service.Log = ctx.Log
	service.ctx = ctx
	return service
}

func (service *Service) Message(message string, codeArgs ...int32) *result.Error {
	service.Log.Info(message)
	return result.Message(message, codeArgs...)
}

func (service *Service) Error(err error, msgarg ...string) *result.Error {
	service.Log.Error("error", zap.Error(err))
	return result.ServerError(err, msgarg...)
}

func ToSystemPage[T any](result *gorme.PageResult[T]) vo.SystemPageListResp[T] {
	var t vo.SystemPageListResp[T]
	t.PageInfo = vo.PageInfo{
		CurrentPage: result.PageNo,
		TotalPage:   result.TotalPage,
		Total:       result.TotalSize,
	}
	t.Items = result.List
	return t
}
