package response

import (
	"fmt"
	"giftcard/library/context/api"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"runtime"
)

const (
	STATUS_SUCESS       = 200
	STATUS_PARAM_ERROR  = 4000
	STATUS_SERVER_ERROR = 5000
)

type Result struct {
	Code    int32  `json:"code"`
	TraceId string `json:"trace_id"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Data    any    `json:"data"`
}

type ErrorMessage struct {
	Code    int32
	Message string
	Detail  string
	error   error
}

func CodeMessage(code int32, message string) *ErrorMessage {
	return &ErrorMessage{
		Code:    code,
		Message: message,
	}
}

func Message(message string) *ErrorMessage {
	return &ErrorMessage{
		Code:    STATUS_PARAM_ERROR,
		Message: message,
	}
}

//第二个参数可选,做一些备注
func Error(err error, msgarg ...string) *ErrorMessage {
	var newError error
	if len(msgarg) > 0 {
		newError = errors.WithMessage(err, msgarg[0])
	} else {
		newError = fmt.Errorf("%w", err)
	}

	detail := err.Error() + traceInfo()
	return &ErrorMessage{
		Code:    STATUS_SERVER_ERROR,
		Message: "服务器错误",
		Detail:  detail,
		error:   newError,
	}
}

//如返回值是error类型时,请使用这个方法
func WrapError(err error, msgarg ...string) error {
	errFuncName := ""
	pc, errFile, errLine, ok := runtime.Caller(1)
	if ok {
		errFuncName = runtime.FuncForPC(pc).Name()
	}

	var newError error
	if len(msgarg) > 0 {
		newError = errors.WithMessage(err, msgarg[0])
	} else {
		newError = fmt.Errorf("%w", err)
	}

	detail := fmt.Sprintf("error:%s at file:%s,line:%d,func:%s", err.Error(), errFile, errLine, errFuncName)
	return errors.WithMessage(newError, detail)
}

func ParamError(ctx *api.Context, message string) {
	response(ctx, STATUS_PARAM_ERROR, message, "", nil)
}

func ServerError(ctx *api.Context, message string) {
	response(ctx, STATUS_SERVER_ERROR, message, "", nil)
}

func ErrorMsg(ctx *api.Context, message *ErrorMessage) {
	response(ctx, message.Code, message.Message, message.Detail, nil)
}

func Data(ctx *api.Context, data interface{}) {
	response(ctx, STATUS_SUCESS, "操作成功", "", data)
}

func Msg(ctx *api.Context, code int32, message string) {
	response(ctx, code, message, "", nil)
}

func Success(ctx *api.Context) {
	response(ctx, STATUS_SUCESS, "操作成功", "", nil)
}

func response(ctx *api.Context, code int32, message string, detail string, data interface{}) {
	if code != STATUS_SUCESS {
		ctx.Log.Error(message, zap.String("detail", detail), zap.String("message", message))
	}
	result := Result{
		Code:    code,
		Message: message,
		Detail:  detail,
		Data:    data,
		TraceId: ctx.TraceId,
	}
	ctx.GinCtx.JSON(STATUS_SUCESS, result)
}

func traceInfo() string {
	errFuncName := ""
	pc, errFile, errLine, ok := runtime.Caller(2)
	if ok {
		errFuncName = runtime.FuncForPC(pc).Name()
	}
	detail := fmt.Sprintf("at file:%s,line:%d,func:%s", errFile, errLine, errFuncName)
	return detail
}
