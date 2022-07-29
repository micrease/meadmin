package router

import (
	"giftcard/application/controller/admin"
	"giftcard/library/context/api"
	"giftcard/system/middleware"
)

func SystemApiRouter(router *api.Router) {
	//管理后台接口
	systemApiGroup := router.Group("/api")
	{
		settingGroup := systemApiGroup.Group("/setting")
		{
			systemConfig := new(admin.SystemConfig)
			settingGroup.GET("/config/getSysConfig", systemConfig.GetSysConfig)
			settingGroup.POST("/config/getConfigByKey", systemConfig.GetConfigByKey)
		}

		systemGroup := systemApiGroup.Group("/system")
		{
			system := new(admin.System)
			systemGroup.POST("/login", system.Login)
			//需要jwt验证的
			systemAuthGroup := systemGroup.Use(middleware.JWTAuth()).Group("/")
			{
				systemGroup.POST("/logout", system.Logout)
				systemAuthGroup.GET("/getInfo", system.GetInfo)
				systemAuthGroup.GET("/user/index", system.PageList)

				dept := admin.SystemDept{}
				systemAuthGroup.GET("/dept/tree", dept.GetListTree)
				systemAuthGroup.GET("/dept/index", dept.Index)

				//role
				role := admin.SystemRole{}
				systemAuthGroup.POST("/role/save", role.Save)
				systemAuthGroup.GET("/role/index", role.Index)
				systemAuthGroup.GET("/role/list", role.List)
				systemAuthGroup.PUT("/role/update/:id", role.Update)
				//menu
				menu := admin.SystemMenu{}
				systemAuthGroup.GET("/menu/index", menu.Index)

				//post
				post := admin.SystemPost{}
				systemAuthGroup.GET("/post/index", post.PageList)
				systemAuthGroup.GET("/post/list", post.PostList)
				systemAuthGroup.POST("/post/save", post.Save)
				systemAuthGroup.PUT("/post/update/:id", post.Update)

				//dataDict/list
				dataDict := admin.SystemDictData{}
				systemAuthGroup.GET("/dataDict/list", dataDict.List)

				//uploadImage
				upload := admin.SystemUploadfile{}
				systemAuthGroup.POST("/uploadImage", upload.Upload)
				systemAuthGroup.GET("/attachment/index", upload.Index)

			}
		}
	}
}
