package router

import (
	"admease/application/controller/admin"
	"admease/library/context/api"
)

func AdminApiRouter(router *api.Router) {
	//管理后台接口
	adminApiGroup := router.Group("/admin")
	{
		//不用验证的接口
		ctrl := new(admin.System)
		adminApiGroup.POST("/login", ctrl.Login)
	}
}
