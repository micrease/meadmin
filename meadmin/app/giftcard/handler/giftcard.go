package handler

import (
	"meadmin/app/giftcard/dto"
	"meadmin/app/giftcard/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

//giftcard
type GiftCard struct {
}

func (this GiftCard) AttrOptions(ctx *api.Context) *result.Result {
	attrs := dto.CardAttrs{}
	attrs.Currencies = []dto.Currency{
		{Code: "CNY", Name: "人民币"},
		{Code: "USD", Name: "美元"},
	}
	attrs.CateList = []dto.CardCate{
		{ID: 1, CateName: "香草VISA"},
		{ID: 2, CateName: "香草Master"},
		{ID: 3, CateName: "Steam"},
	}

	attrs.CardPrefixs = []dto.CardPrefix{
		{ID: 1, Name: "1111"},
		{ID: 2, Name: "6666"},
		{ID: 3, Name: "5555"},
		{ID: 4, Name: "8888"},
	}
	return result.Success(attrs)
}

//保存信息
func (this GiftCard) Save(ctx *api.Context) *result.Result {
	var req dto.MerchantSaveReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewMerchant(ctx).Save(req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}
