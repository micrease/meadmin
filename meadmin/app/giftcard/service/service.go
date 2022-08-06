package service

import (
	"github.com/micrease/gorme"
	"go.uber.org/zap"
	"meadmin/app/system/dto"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/logger"
)

type Service struct {
	ctx     *api.Context
	TraceId string
	Log     logger.Log
}

func (s *Service) Init(ctx *api.Context) *Service {
	s.TraceId = ctx.TraceId
	s.Log = ctx.Log
	s.ctx = ctx
	return s
}

func (s *Service) Message(message string, codeArgs ...int32) *result.Error {
	s.Log.Info(message)
	return result.Message(message, codeArgs...)
}

func (s *Service) Error(err error, msgarg ...string) *result.Error {
	s.Log.Error("error", zap.Error(err))
	return result.ServerError(err, msgarg...)
}

func ToSystemPage[T any](result *gorme.PageResult[T]) dto.SystemPageListResp[T] {
	var t dto.SystemPageListResp[T]
	t.PageInfo = dto.PageInfo{
		CurrentPage: result.PageNo,
		TotalPage:   result.TotalPage,
		Total:       result.TotalSize,
	}
	t.Items = result.List
	return t
}
