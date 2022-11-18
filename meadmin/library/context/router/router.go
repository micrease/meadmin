package router

import (
	"github.com/gin-gonic/gin"
	"meadmin/library/context/api"
	"meadmin/library/rpc"
	"reflect"
	"runtime"
	"strings"
)

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

// Router相关
func NewRouter(engin *gin.Engine) *Router {
	ctx := new(Router)
	ctx.GinEngin = engin
	return ctx
}

// 新版接口
func (router *RouterGroup) GET(relativePath string, handler api.ResultHandleFunc) {
	router.GinRouterGroup.GET(relativePath, api.ResultHandle(handler))
	router.RegiserRpcRouter(router.GinRouterGroup.BasePath()+relativePath, handler)
}

func (router *RouterGroup) Any(relativePath string, handler api.ResultHandleFunc) {
	router.GinRouterGroup.Any(relativePath, api.ResultHandle(handler))
	router.RegiserRpcRouter(router.GinRouterGroup.BasePath()+relativePath, handler)
}

func (router *RouterGroup) RegiserRpcRouter(relativePath string, handler api.ResultHandleFunc) {
	relativePath = strings.ReplaceAll(relativePath, "//", "/")
	fullName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

	fname := ""
	if fullName != "" {
		arr := strings.Split(fullName, ".")
		if len(arr) == 2 {
			fname = arr[1]
		} else if len(arr) == 3 {
			fname = arr[1] + "." + arr[2]
		}

		fname = strings.ReplaceAll(fname, "-fm", "")
		fname = strings.ReplaceAll(fname, "*", "")
		fname = strings.ReplaceAll(fname, "(", "")
		fname = strings.ReplaceAll(fname, ")", "")
	}

	rpc.GetRpcServer().Handle(relativePath, handler, "")
	rpc.GetRpcServer().Handle(fname, handler, "")
}

func (router *RouterGroup) POST(relativePath string, handler api.ResultHandleFunc) {

	router.RegiserRpcRouter(router.GinRouterGroup.BasePath()+relativePath, handler)
	router.GinRouterGroup.POST(relativePath, api.ResultHandle(handler))
}

func (router *RouterGroup) PUT(relativePath string, handler api.ResultHandleFunc) {
	router.GinRouterGroup.PUT(relativePath, api.ResultHandle(handler))
	router.RegiserRpcRouter(router.GinRouterGroup.BasePath()+relativePath, handler)
}

func (router *RouterGroup) DELETE(relativePath string, handler api.ResultHandleFunc) {
	router.GinRouterGroup.DELETE(relativePath, api.ResultHandle(handler))
	router.RegiserRpcRouter(router.GinRouterGroup.BasePath()+relativePath, handler)
}

func (router *RouterGroup) HEAD(relativePath string, handler api.ResultHandleFunc) {
	router.GinRouterGroup.HEAD(relativePath, api.ResultHandle(handler))
}

func (router *RouterGroup) OPTIONS(relativePath string, handler api.ResultHandleFunc) {
	router.GinRouterGroup.OPTIONS(relativePath, api.ResultHandle(handler))
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

// middleware
func (router *Router) Use(middleware ...gin.HandlerFunc) *Router {
	router.GinEngin.Use(middleware...)
	return router
}

// group
func (router *Router) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	routerGroup := new(RouterGroup)
	ginGroup := router.GinEngin.Group(relativePath, handlers...)
	routerGroup.Router = router
	routerGroup.GinRouterGroup = ginGroup
	return routerGroup
}
