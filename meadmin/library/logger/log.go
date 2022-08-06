package logger

import (
	"go.uber.org/zap"
)

const TraceId = "TraceId"

type Log struct {
	TraceId string
}

//Log部分
func (log Log) Debug(msg string, fields ...zap.Field) {
	fields = log.appendTraceId(fields...)
	Logger.Debug(msg, fields...)
}

func (log Log) Info(msg string, fields ...zap.Field) {
	fields = log.appendTraceId(fields...)
	Logger.Info(msg, fields...)
}

func (log Log) Warn(msg string, fields ...zap.Field) {
	fields = log.appendTraceId(fields...)
	Logger.Warn(msg, fields...)
}

func (log Log) Error(msg string, fields ...zap.Field) {
	fields = log.appendTraceId(fields...)
	Logger.Error(msg, fields...)
}

//在log中植入请求数据
func (log Log) appendTraceId(fields ...zap.Field) []zap.Field {
	fields = append(fields, zap.String(TraceId, log.TraceId))
	return fields
}
