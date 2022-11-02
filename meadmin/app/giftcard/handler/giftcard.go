package handler

import (
	"github.com/spf13/cast"
	"meadmin/app/giftcard/dto"
	"meadmin/app/giftcard/service"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
	"strings"
)

// GiftCard
// @Description:
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
	var req dto.GiftCardSaveReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewGiftCard(ctx).Save(req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

func (this GiftCard) PageList(ctx *api.Context) *result.Result {
	var req dto.GiftCardPageListReq
	validate.BindWithPanic(ctx, &req)
	list, err := service.NewGiftCard(ctx).PageList(req)
	if err != nil {
		return result.ErrorMessage(err, "查询失败")
	}
	return result.Success(list)
}

//修改状态
func (this GiftCard) ChangeStatus(ctx *api.Context) *result.Result {
	var req dto.GiftCardStatusReq
	validate.BindWithPanic(ctx, &req)
	err := service.NewGiftCard(ctx).ChangeStatus(req)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}

//Detail
func (this GiftCard) Detail(ctx *api.Context) *result.Result {
	reqId := ctx.GinCtx.Param("id")
	id := cast.ToUint(reqId)
	if id <= 0 {
		return result.Success()
	}

	model, err := service.NewGiftCard(ctx).Detail(id)
	if err != nil {
		return result.ErrorMessage(err)
	}
	return result.Success(model)
}

// Delete
// @Description: 删除礼品卡
// @receiver g
// @param ctx
// @return *result.Result
func (g GiftCard) Delete(ctx *api.Context) *result.Result {
	id := ctx.GinCtx.Param("id")
	ids := strings.Split(id, ",")
	err := service.NewGiftCard(ctx).Delete(ids)
	if err != nil {
		return result.ErrorResult(err)
	}
	return result.Success()
}
