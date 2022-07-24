package router

import (
	"giftcard/library/context/api"
	"giftcard/system/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

//这里是一个演示中间件
func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
	}
}

func InitGinRouter() *gin.Engine {
	//初始化数据
	gin.SetMode(gin.DebugMode)
	ginRouter := gin.Default()
	//这里我们用自己的context对gin进行包装。
	//目的是我们可以动态向请求体中植入自定义的变量或方法,来扩展业务逻辑中的通用方法
	//不同于middleware仅作用于请求上下文, context作用于框架上下文
	//如果需要使用gin的原生方法,可以router.GinEngin获取,
	//但是建议你通过包装一层。尽量使用library下的context访问
	//原生方法:
	//router.GinEngin.GET("/gin",nil)
	router := api.NewRouter(ginRouter)

	//中间件
	router.Use(middleware.Recover())
	router.Use(middleware.Cors())
	router.Use(middleware.RequestLog())

	SystemApiRouter(router)
	AdminApiRouter(router)
	staticRouter(router)
	return router.GinEngin
}

func staticRouter(router *api.Router) {
	router.GinEngin.LoadHTMLGlob("templates/**/*")
	staticGroup := router.GinEngin.Group("/")
	staticGroup.GET("/mock", func(ginCtx *gin.Context) {
		ginCtx.HTML(http.StatusOK, "task/index.html", gin.H{
			"title": "posts/index",
		})
	})
}
