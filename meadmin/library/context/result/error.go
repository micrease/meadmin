package result

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

//Error可以在各层中使用,Result只建议在Controller中使用.
//Error表示系统中存在错误,错误包含业务错误和系统错误。业务错误比如:库存不足。系统错误如:数据库字段不存在
//Result表示返回的是一种结果,结果也包含错误,因此错误可以转化为结果
type ErrorMessage struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	error   error  `json:"error"`
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
