package result

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

type Error struct {
	Code    int32  `json:"code"`    //业务状态码,200表示请求成功,其它表示错误
	Message string `json:"message"` //业务提示消息
	Detail  string `json:"detail"`  //错误详情
}

func ParamError(err error, msgarg ...string) *Error {
	msg := err.Error()
	if len(msgarg) > 0 {
		msg = msgarg[0]
	}
	detail := err.Error() + traceInfo()
	return &Error{
		Code:    STATUS_PARAM_ERROR,
		Message: msg,
		Detail:  detail,
	}
}

func ServerError(err error, msgarg ...string) *Error {
	msg := err.Error()
	if len(msgarg) > 0 {
		msg = msgarg[0]
	}

	detail := err.Error() + traceInfo()
	return &Error{
		Code:    STATUS_SERVER_ERROR,
		Message: msg,
		Detail:  detail,
	}
}

func Message(message string, codeArgs ...int32) *Error {
	var code int32 = STATUS_PARAM_ERROR
	if len(codeArgs) > 0 {
		code = codeArgs[0]
	}
	return &Error{
		Code:    code,
		Message: message,
		Detail:  message,
	}
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
