package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"math/rand"
	"meadmin/library/apidoc"
	"meadmin/library/astparser"
	"meadmin/library/context/result"
	"meadmin/library/logger"
	"meadmin/library/rpc/protocol"
	"meadmin/library/strings"
	"meadmin/system/config"
	"meadmin/system/consts"
)

type ResultHandleFunc func(ctx *Context) *result.Result
type Context struct {
	TraceId      string
	ProjectPath  string
	ApiDoc       astparser.ApiDoc //文档生成信息
	GinCtx       *gin.Context     //gin的Context,只有在gin接口调用时有值
	RpcCtx       context.Context  //rpc 的Context,只有在rpc调用时有值
	RpcArgs      *protocol.Args   //rpc 的参数,只有在rpc调用时有值
	Log          logger.Log
	JwtClaimData consts.JwtClaimData
	error        error
	IsRpcRequest bool //是否rpc请求
}

type HandleFunc func(ctx *Context)

// 这里我们用自己的context包对gin进行包装。
// 目的是我们可以动态向请求体中植入自定义的变量或方法,来扩展业务逻辑中的通用方法
// 不同于middleware作用于请求上下文, context作用于框架上下文
// 如果需要使用gin的原生方法,可以router.GinEngin获取,
// 但是建议你通过包装一层。尽量使用library下的context访问
// 原生方法:
// router.GinEngin.GET("/gin",nil)

// 与gin兼容
func Handle(fn HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := NewContext(c)
		//这里调用contrller的方法,因此你可以在这个方法前后进行各种前置,后置操作
		fn(ctx)
	}
}

// 带返回值的ctrl
func ResultHandle(fn ResultHandleFunc) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctx := NewContext(ginCtx)
		//这里调用contrller的方法,因此你可以在这个方法前后进行各种前置,后置操作
		result := fn(ctx)
		result.TraceId = ctx.TraceId
		//如果开启了文档,并且返回正确
		if ctx.ApiDoc.Enable && result.Code == 200 {
			GenDoc(ctx, fn, result)
		}
		ginCtx.JSON(200, result)
	}
}

// 每个请求一个Context
func NewContext(ginCtx *gin.Context) *Context {
	ctx := new(Context)
	ctx.GinCtx = ginCtx

	conf := config.GetConfig()
	if conf.IsDocEnable() {
		ctx.ApiDoc.Enable = true
	}

	ctx.ProjectPath = conf.ProjectPath
	traceId := ginCtx.GetHeader(logger.TraceId)
	if len(traceId) <= 0 {
		traceId = ginCtx.GetHeader("RequestId")
	}

	if len(traceId) <= 0 {
		ctxTraceId, ok := ginCtx.Get(logger.TraceId)
		if ok {
			traceId = cast.ToString(ctxTraceId)
		}
	}

	if len(traceId) <= 0 {
		traceId = fmt.Sprintf("R%d", rand.Int63())
		ctx.GinCtx.Header(logger.TraceId, traceId)
	}

	ctx.TraceId = traceId
	ctx.Log.TraceId = ctx.TraceId

	//解析token中的附加信息
	jwtClaimData, exist := ginCtx.Get(consts.JwtClaimDataKey)
	if exist {
		if jwtData, ok := jwtClaimData.(consts.JwtClaimData); ok {
			ctx.JwtClaimData = jwtData
		}
	}
	return ctx
}

// 每个请求一个Context
func NewRPCContext(rpcCtx context.Context, args *protocol.Args) *Context {
	ctx := new(Context)
	conf := config.GetConfig()
	if conf.IsDocEnable() {
		ctx.ApiDoc.Enable = true
	}

	ctx.IsRpcRequest = true
	ctx.RpcCtx = rpcCtx
	ctx.RpcArgs = args

	traceId := ""
	ctx.ProjectPath = conf.ProjectPath
	if len(traceId) <= 0 {
		traceId = fmt.Sprintf("R%d", rand.Int63())
	}

	ctx.TraceId = traceId
	ctx.Log.TraceId = ctx.TraceId
	return ctx
}

func GenDoc(ctx *Context, fn ResultHandleFunc, result *result.Result) {
	var err error
	ctx.ApiDoc.TraceId = ctx.TraceId
	ctx.ApiDoc.Method = ctx.GinCtx.Request.Method
	ctx.ApiDoc.ApiName = ctx.GinCtx.FullPath()
	ctx.ApiDoc.ResultAst, err = astparser.ParseStruct(ctx.ProjectPath, result)
	if err != nil {
		ctx.Log.Error("astparser.ParseStruct Error", zap.Error(err))
	}
	ctx.ApiDoc.FuncAst, err = astparser.ParseFunc(ctx.ProjectPath, fn)
	if err != nil {
		ctx.Log.Error("astparser.ParseFunc Error", zap.Error(err))
	}

	resultJson, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		ctx.Log.Error("json.Marshal(result) Error", zap.Error(err))
	}

	resultStr := string(resultJson)
	//数据太长了。截取一部分写在文档中,作为示例数据
	maxLength := 2048
	if len(resultStr) > maxLength {
		ctx.ApiDoc.ResultExample = strings.Substr(resultStr, 0, maxLength)
		ctx.ApiDoc.ResultExample += "........更多示例数据不再展示"
	} else {
		ctx.ApiDoc.ResultExample = resultStr
	}
	apidoc.SaveApi(&ctx.ApiDoc)
}
