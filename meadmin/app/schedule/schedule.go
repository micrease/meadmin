package schedule

import (
	"admease/library/logger"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"runtime"
	"runtime/debug"
)

/**
 * 每秒执行一次，每次执行时长为2秒。start每秒都会执行，并不是串行隔5秒后才运行。
 */
func goroutineWatcher() {
	num := runtime.NumGoroutine()
	logger.Info("NumGoroutine", zap.Int("num", num))
}

func Start() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("schedule defer Error", zap.Any("err", err))
			debug.PrintStack()
		}
	}()

	watcher := cron.New(cron.WithSeconds())
	watcher.AddFunc("@every 1m", goroutineWatcher)
	//构建用户索引
	watcher.Start()
}

func Test() {
	c := cron.New()
	c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour, starting an hour from now") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
	select {}
}
