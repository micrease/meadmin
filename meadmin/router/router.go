package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meadmin/library/context/router"
	"meadmin/system/config"
	"meadmin/system/middleware"
	"net/http"
)

// 这里是一个演示中间件
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
	router := router.NewRouter(ginRouter)
	//原生方法:
	//router.GinEngin.GET("/gin", func(context *gin.Context) {
	//})

	//中间件
	router.Use(middleware.Recover())
	router.Use(middleware.Cors())
	router.Use(middleware.RequestLog())

	systemGroup := SystemApiRouter(router)
	GiftCardRouter(systemGroup)
	AdminApiRouter(router)
	staticRouter(router)
	return router.GinEngin
}

func staticRouter(router *router.Router) {
	router.GinEngin.LoadHTMLGlob("templates/*.html")
	staticGroup := router.GinEngin.Group("/")
	staticGroup.GET("/index", func(ginCtx *gin.Context) {
		ginCtx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
		})
	})

	conf := config.GetConfig()
	path := conf.ProjectPath + "/runtime/upload/"
	fmt.Println(path)
	router.GinEngin.Static("/static/", path)
}
