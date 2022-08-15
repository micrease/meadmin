package router

import (
	"meadmin/app/giftcard/handler"
	"meadmin/library/context/api"
	"meadmin/system/middleware"
)

func GiftCardRouter(systemApiGroup *api.RouterGroup) {
	giftcardGroup := systemApiGroup.Use(middleware.JWTAuth()).Group("/giftcard")
	{
		merchant := handler.Merchant{}
		giftcardGroup.POST("/merchant/save", merchant.Save)
		giftcardGroup.GET("/merchant/index", merchant.PageList)
		giftcardGroup.PUT("/merchant/change_status", merchant.ChangeStatus)
		giftcardGroup.GET("/merchant/detail/:id", merchant.Detail)

		giftcard := handler.GiftCard{}
		giftcardGroup.GET("/attr_options", giftcard.AttrOptions)
	}
}
