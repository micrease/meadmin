package router

import (
	"meadmin/app/system/handler"
	"meadmin/library/context/api"
	"meadmin/system/middleware"
)

func SystemApiRouter(router *api.Router) *api.RouterGroup {
	//管理后台接口
	systemApiGroup := router.Group("/api")
	{
		settingGroup := systemApiGroup.Group("/setting")
		{
			systemConfig := handler.SystemConfig{}
			settingGroup.GET("/config/getSysConfig", systemConfig.GetSysConfig)
			settingGroup.POST("/config/getConfigByKey", systemConfig.GetConfigByKey)
		}

		systemGroup := systemApiGroup.Group("/system")
		{
			system := handler.System{}
			systemGroup.POST("/login", system.Login)
			//需要jwt验证的
			systemAuthGroup := systemGroup.Use(middleware.JWTAuth()).Group("/")
			{
				systemGroup.POST("/logout", system.Logout)
				systemAuthGroup.GET("/getInfo", system.GetInfo)
				systemAuthGroup.GET("/user/index", system.PageList)
				//systemAuthGroup.GET("/user/read/:id", system.ReadInfo)

				dept := handler.SystemDept{}
				systemAuthGroup.GET("/dept/tree", dept.GetListTree)
				systemAuthGroup.GET("/dept/index", dept.Index)

				//role
				role := handler.SystemRole{}
				systemAuthGroup.POST("/role/save", role.Save)
				systemAuthGroup.GET("/role/index", role.Index)
				systemAuthGroup.GET("/role/list", role.List)
				systemAuthGroup.PUT("/role/update/:id", role.Update)
				//menu
				menu := handler.SystemMenu{}
				systemAuthGroup.GET("/menu/index", menu.Index)

				//post
				post := handler.SystemPost{}
				systemAuthGroup.GET("/post/index", post.PageList)
				systemAuthGroup.GET("/post/list", post.PostList)
				systemAuthGroup.POST("/post/save", post.Save)
				systemAuthGroup.PUT("/post/update/:id", post.Update)

				//dataDict/list
				dataDict := handler.SystemDictData{}
				systemAuthGroup.GET("/dataDict/list", dataDict.List)

				//uploadImage
				upload := handler.SystemUploadfile{}
				systemAuthGroup.POST("/uploadImage", upload.Upload)
				systemAuthGroup.GET("/attachment/index", upload.Index)
			}
		}
	}
	return systemApiGroup
}
