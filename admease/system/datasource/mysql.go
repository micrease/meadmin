package datasource

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

//管理多库连接,读写分离或多库,微服务中不建议连接多个库，如果需要的话可以在这里扩展
type DBManager struct {
	master *gorm.DB //主库
	slave  *gorm.DB //从库，读库
}

var dbManager *DBManager

func InitDatabase(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalln("连接数据库失败")
	}
	dbManager = new(DBManager)
	dbManager.master = db
}

func GetDB() *gorm.DB {
	return dbManager.master
}

func GetDBSlave() *gorm.DB {
	return dbManager.slave
}
