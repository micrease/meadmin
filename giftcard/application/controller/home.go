package controller

import (
	"errors"
	"giftcard/library/context/api"
	"giftcard/library/context/result"
	"giftcard/library/response"
	"time"
)

type Home struct {
	BaseController
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
