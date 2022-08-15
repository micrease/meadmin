package handler

import (
	"github.com/spf13/cast"
	"meadmin/app/giftcard/dto"
	"meadmin/app/giftcard/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

//商户管理
type Merchant struct {
}

//保存信息
func (this Merchant) Save(ctx *api.Context) *result.Result {
	var req dto.MerchantSaveReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewMerchant(ctx).Save(req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

//修改状态
func (this Merchant) ChangeStatus(ctx *api.Context) *result.Result {
	var req dto.MerchantStatusReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewMerchant(ctx).ChangeStatus(req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

//Detail
func (this Merchant) Detail(ctx *api.Context) *result.Result {
	reqId := ctx.GinCtx.Param("id")
	id := cast.ToUint(reqId)
	if id <= 0 {
		return result.Success()
	}

	model, err := service.NewMerchant(ctx).Detail(id)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(model)
}

//获取列表
func (this Merchant) PageList(ctx *api.Context) *result.Result {
	var req dto.MerchantPageListReq
	validate.BindWithPanic(ctx, &req)

	list, err := service.NewMerchant(ctx).PageList(req)
	if err != nil {
		return result.ErrorMessage(err, "查询失败")
	}
	return result.Success(list)
}
