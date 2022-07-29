package handler

import (
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/response"
	"errors"
	"time"
)

type Home struct {
	Handler
}

/**
 * 获取查询列表
 */
func (ctrl *Home) Index(ctx *api.Context) *result.Result {
	ctx.Log.Info("Check Health")
	err := result.Error(errors.New("ffff"), "test error")
	return result.Success(err)
}

func (ctrl *Home) Test(ctx *api.Context) {
	response.Msg(ctx, 200, time.Now().String())
}
