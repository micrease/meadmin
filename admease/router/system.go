package router

import (
	"admease/app/system/handler"
	"admease/library/context/api"
	"admease/system/middleware"
)

func SystemApiRouter(router *api.Router) {
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

				SystemUser := handler.SystemUser{}
				systemAuthGroup.GET("/user/read/:id", SystemUser.ReadInfo)
				systemAuthGroup.POST("/user/clearCache", SystemUser.ClearCache)
				systemAuthGroup.PUT("/user/changeStatus", SystemUser.ChangeStatus)
				systemAuthGroup.POST("/user/save", SystemUser.SaveUser)
				systemAuthGroup.DELETE("/user/delete/:id", SystemUser.DeleteUser)
				systemAuthGroup.PUT("/user/initUserPassword/:id",SystemUser.InitUserPassword)
				systemAuthGroup.POST("/user/setHomePage", SystemUser.SetHomePage)
				systemAuthGroup.PUT("/user/update/:id", SystemUser.Update)

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
}
