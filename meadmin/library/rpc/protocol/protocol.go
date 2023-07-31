package protocol

type Args struct {
	Data        any               //任意类型
	HandlerName string            //回调方法名
	Headers     map[string]string //类似http一样的通用header请求头
}

type Reply struct {
	Data any
}

type RpcResult[T any] struct {
	Code    int32  `json:"code"`     //业务状态码,200表示请求成功,其它表示错误
	Message string `json:"message"`  //业务提示消息
	Detail  string `json:"detail"`   //错误详情
	TraceId string `json:"trace_id"` //请求链路ID,可用来方便查询log
	Success bool   `json:"success"`  //是否成功
	Data    T      `json:"data"`     //返回数据体
}
