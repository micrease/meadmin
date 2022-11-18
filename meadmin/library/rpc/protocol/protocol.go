package protocol

type Config struct {
	Network        string //tcp,http
	Listen         string //ip:port
	BasePath       string
	ServiceName    string
	DiscoveryAddrs []string
}

type Args struct {
	Data        any               //任意类型
	HandlerName string            //回调方法名
	Headers     map[string]string //类似http一样的通用header请求头
}

type Reply struct {
	Data any
}
