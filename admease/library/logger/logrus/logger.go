package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"math/rand"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultLogPath = "./runtime/logs"
)

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("01-02 15:04:05.000")
	var file string
	var line int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}

	data := ""
	if len(entry.Data) > 0 {
		json, _ := json.Marshal(entry.Data)
		data = (string(json))
	}

	msg := fmt.Sprintf("%s[Goroutine:%d][%s]%s,%s,at:%s:%d \n",
		timestamp, getGID(), strings.ToUpper(entry.Level.String()), entry.Message, data, file, line)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

/**
日志轮转相关函数:
WithLinkName 为最新的日志建立软连接
WithRotationTime 设置日志分割的时间，隔多久分割一次
WithMaxAge 和 WithRotationCount二者只能设置一个
WithMaxAge 设置文件清理前的最长保存时间
WithRotationCount 设置文件清理前最多保存的个数
*/
func InitLog() {
	path := DefaultLogPath + "/run.log"
	//下面配置日志每隔 1 小时轮转一个新文件，保留最近 24小时的日志文件，多余的自动清理掉。
	//日志保留时间
	logKeepMaxHours := 24
	//每小时一个文件
	everyLogFileDurationHours := 24
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(logKeepMaxHours)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(everyLogFileDurationHours)*time.Hour),
	)
	logrus.SetOutput(writer)
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(new(LogFormatter))
	/*
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	*/
}

type Map map[string]interface{}

func Data(data Map) *logrus.Entry {
	return logrus.WithFields(logrus.Fields(data))
}

func Info(args ...interface{}) {
	logrus.Info(args)
}

func RandInfo(args ...interface{}) {
	n := rand.Int31n(10)
	if n == 1 {
		logrus.Info("[RandInfo]", args)
	}
}

func Error(args ...interface{}) {
	logrus.Error(args)
}
