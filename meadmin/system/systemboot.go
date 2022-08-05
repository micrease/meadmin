package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meadmin/app/schedule"
	"meadmin/library/context/api"
	"meadmin/library/logger"
	"meadmin/library/migrate"
	redisClient "meadmin/library/redis"
	"meadmin/library/validate"
	"meadmin/router"
	"meadmin/system/config"
	"meadmin/system/datasource"
	"meadmin/system/middleware"
)

//1,初始化服务
func ServerInit(configPath string) {
	fmt.Println("configPath:", configPath)
	//1,解析配置文件
	conf := config.InitConfig(configPath)
	//2,Log初始化
	logger.InitLog()
	//3,加载参数验证扩展
	validate.RegisterExtension()
	//4,数据库初始化
	datasource.InitDatabase(conf.Database.Dsn)
	//5,redis初始化
	redisClient.Connect(conf.Redis.Addr)
	//6,启动时自动执行migrate
	migrate.Install()
	//7,业务初始化
	appInit()
}

//2,业务应用初始化
func appInit() {

}

//3,启动服务
func ServerRun() {
	conf := config.GetConfig()
	fmt.Println("Server Run Start")
	logger.Info("Server Run Start")
	//1,Gin框架初始化
	route := router.InitGinRouter()
	//2,启动定时任务
	schedule.Start()
	if route.Run(":"+conf.Port) != nil {
		logger.Error("Server Run Error")
	}
}

//测试用例的入口
func SetupTestRouter() *gin.Engine {
	//设为release,要不然输出的东西太多，影响视线
	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()
	ctxRouter := api.NewRouter(ginRouter)
	ctxRouter.Use(middleware.Recover())
	ctxRouter.Use(middleware.Cors())
	ctxRouter.Use(middleware.RequestLog())

	router.AdminApiRouter(ctxRouter)
	return ginRouter
}
