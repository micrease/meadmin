package api

import "github.com/gin-gonic/gin"

//这里我们用自己的context包对gin进行包装。
//目的是我们可以动态向请求体中植入自定义的变量或方法,来扩展业务逻辑中的通用方法
//不同于middleware仅作用于请求上下文, context作用于框架上下文
//如果需要使用gin的原生方法,可以router.GinEngin获取,
//但是建议你通过包装一层。尽量使用library下的context访问
//原生方法:
//router.GinEngin.GET("/gin",nil)

type RouterGroup struct {
	GinRouterGroup *gin.RouterGroup
	Router         *Router
}

type Router struct {
	GinEngin    *gin.Engine
	RouterGroup *RouterGroup
}

//Router相关
func NewRouter(engin *gin.Engine) *Router {
	ctx := new(Router)
	ctx.GinEngin = engin
	return ctx
}

//新版接口
func (router *RouterGroup) GET(relativePath string, handler ResultHandleFunc) {
	router.GinRouterGroup.GET(relativePath, ResultHandle(handler))
}

func (router *RouterGroup) Any(relativePath string, handler ResultHandleFunc) {
	router.GinRouterGroup.Any(relativePath, ResultHandle(handler))
}

func (router *RouterGroup) POST(relativePath string, handler ResultHandleFunc) {
	router.GinRouterGroup.POST(relativePath, ResultHandle(handler))
}

func (router *RouterGroup) PUT(relativePath string, handler ResultHandleFunc) {
	router.GinRouterGroup.PUT(relativePath, ResultHandle(handler))
}

func (router *RouterGroup) DELETE(relativePath string, handler ResultHandleFunc) {
	router.GinRouterGroup.DELETE(relativePath, ResultHandle(handler))
}

func (router *RouterGroup) HEAD(relativePath string, handler ResultHandleFunc) {
	router.GinRouterGroup.HEAD(relativePath, ResultHandle(handler))
}

func (router *RouterGroup) OPTIONS(relativePath string, handler ResultHandleFunc) {
	router.GinRouterGroup.OPTIONS(relativePath, ResultHandle(handler))
}

func (router *RouterGroup) Use(middleware ...gin.HandlerFunc) *RouterGroup {
	router.GinRouterGroup.Use(middleware...)
	return router
}

func (router *RouterGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	ginGroup := router.GinRouterGroup.Group(relativePath, handlers...)
	routerGroup := new(RouterGroup)
	routerGroup.GinRouterGroup = ginGroup
	return routerGroup
}

//middleware
func (router *Router) Use(middleware ...gin.HandlerFunc) *Router {
	router.GinEngin.Use(middleware...)
	return router
}

//group
func (router *Router) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	routerGroup := new(RouterGroup)
	ginGroup := router.GinEngin.Group(relativePath, handlers...)
	routerGroup.Router = router
	routerGroup.GinRouterGroup = ginGroup
	return routerGroup
}
