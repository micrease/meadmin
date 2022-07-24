package main

import (
	"flag"
	"giftcard/system"
	"giftcard/system/config"
)

var (
	//指定配置文件运行环境,默认开发环境
	env = flag.String("env", "dev", "dev/test/prod")
)

func main() {
	//1,解析命令运行参数
	flag.Parse()
	//2,获取配置文件地址
	configPath := config.GetConfigPath(*env)
	//初始化
	system.ServerInit(configPath)
	//运行
	system.ServerRun()
}
