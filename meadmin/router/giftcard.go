package router

import (
	"meadmin/app/giftcard/handler"
	"meadmin/library/context/router"
	"meadmin/system/middleware"
)

func GiftCardRouter(systemApiGroup *router.RouterGroup) {
	giftcardGroup := systemApiGroup.Use(middleware.JWTAuth()).Group("/giftcard")
	{
		merchant := handler.Merchant{}
		giftcardGroup.POST("/merchant/save", merchant.Save)
		giftcardGroup.GET("/merchant/index", merchant.PageList)
		giftcardGroup.PUT("/merchant/change_status", merchant.ChangeStatus)
		giftcardGroup.GET("/merchant/detail/:id", merchant.Detail)

		giftcard := handler.GiftCard{}
		giftcardGroup.GET("/attr_options", giftcard.AttrOptions)
		giftcardGroup.GET("/index", giftcard.PageList)
		giftcardGroup.POST("/save", giftcard.Save)
		giftcardGroup.PUT("/change_status", giftcard.ChangeStatus)
		giftcardGroup.GET("/detail/:id", giftcard.Detail)
		giftcardGroup.DELETE("/delete/:id", giftcard.Delete)

	}
}
