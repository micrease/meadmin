package migrate

import (
	"admease/library/logger"
	"admease/system/config"
	"database/sql"
	"github.com/gobuffalo/packr/v2"
	"github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

func Install() {
	conf := config.GetConfig()
	ds := conf.Database.Dsn
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "../../databases/migrations"),
	}

	db, err := sql.Open("mysql", ds)
	if err != nil {
		logger.Error("Migrate Connect Mysql Error:" + err.Error())
		panic("Migrate Connect Mysql Error:" + err.Error())
	}

	migrate.SetIgnoreUnknown(true)

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		logger.Error("Migrate Exec Up Warn:" + err.Error())
	}
	logger.Info("Applied Up migrations!", zap.Int("files", n))

	/*
		n, err = migrate.Exec(db, "mysql", migrations, migrate.Down)
		if err != nil {
			logger.Warn("Migrate Exec Down Warn:" + err.Error())
		}
		logger.Info("Applied Down migrations!", zap.Int("files", n))
	*/
}
