package handler

import (
	"errors"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/response"
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
	err := result.ErrorMessage(errors.New("ffff"), "test error")
	return result.Success(err)
}

func (ctrl *Home) Test(ctx *api.Context) {
	response.Msg(ctx, 200, time.Now().String())
}
