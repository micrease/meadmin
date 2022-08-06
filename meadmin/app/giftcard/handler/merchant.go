package handler

import (
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

//获取列表
func (this Merchant) List(ctx *api.Context) *result.Result {
	var req dto.MerchantPageListReq
	validate.BindWithPanic(ctx, &req)

	list, err := service.NewMerchant(ctx).PageList(req)
	if err != nil {
		return result.ErrorMessage(err, "查询失败")
	}
	return result.Success(list)
}
