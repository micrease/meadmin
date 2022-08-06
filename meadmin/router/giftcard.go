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
	}
}
