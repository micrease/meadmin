
# 通用后台系统

## 配置说明

```
配置文件在resources目录中
dev,test,prod环境分别对应配置文件:
config-dev.ini
config-test.ini
config-prod.ini
```

## 布署说明

1,开发语言:go,需要安装golang1.18版本  
2,部署相关命令

```
run,start,restart命令需要指定--env参数,以激活不同环境对应的配置文件

sh service.sh   #查看相关命令
sh service.sh run --env=prod  #编译并运行,等效于(build->stop->start),推荐使用,env可选参数有dev,test,prod
sh service.sh build  #编译
sh service.sh start --env=prod #启动
sh service.sh stop  #停止
sh service.sh restart  --env=prod #重启,等效于(stop->start)
```

## 正式环境布署

```
#编译并重启服务
sh service.sh run --env=prod
```

## 测试环境布署

```
#编译并重启服务
sh service.sh run --env=test
```

## 开发说明

1,由于在代码中使用了范型,Golang最低版本1.18,Goland最低2022.1

### Migrate

migrate会在服务启动时自动执行,文件放在databases/migrations目录,文件格式参见:  
https://github.com/rubenv/sql-migrate  
如果执行过程中某一个sql出错,整批的sql都不会成功，需要修复后重新执行。符合事务特点

### Yapi自动生成文档

1,需要在配置文件中增加:

```
doc_enable = true
```

2,自动生成文档时,需要给接口名(controller层的方法上面),请求对象和返回对象字段增加普通注释(多行或单行,不限制格式)

### 相关库

* web路由框架gin  
* orm框架gorme:  
* 时间库carbon  
* 类型转换库cast  
* 前端UI框架SCUI: https://lolicode.gitee.io/scui-doc/  
* mineadmin框架:https://www.mineadmin.com/