package result

//Error可以在各层中使用,Result只建议在Controller中使用.
//Error表示系统中存在错误,错误包含业务错误和系统错误。业务错误比如:库存不足。系统错误如:数据库字段不存在
//Result表示返回的是一种结果,结果也包含错误,因此错误可以转化为结果

//返回结果结构
type Result struct {
	Code    int32  `json:"code"`     //业务状态码,200表示请求成功,其它表示错误
	Message string `json:"message"`  //业务提示消息
	Detail  string `json:"detail"`   //错误详情
	TraceId string `json:"trace_id"` //请求链路ID,可用来方便查询log
	Success bool   `json:"success"`  //是否成功
	Data    any    `json:"data"`     //返回数据体
}

func Success(dataArgs ...any) *Result {
	var data any
	if len(dataArgs) > 0 {
		data = dataArgs[0]
	}

	return &Result{
		Code:    STATUS_SUCESS,
		Message: "成功",
		Success: true,
		Data:    data,
	}
}

func ErrorMessage(err error, msgarg ...string) *Result {
	msg := err.Error()
	if len(msgarg) > 0 {
		msg = msgarg[0]
	}
	return &Result{
		Code:    STATUS_SERVER_ERROR,
		Message: msg,
		Detail:  err.Error(),
		Success: false,
	}
}

func ErrorResult(resultError *Error) *Result {
	return &Result{
		Code:    resultError.Code,
		Message: resultError.Message,
		Success: false,
	}
}
