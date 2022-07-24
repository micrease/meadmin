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
	fields = log.AppendFields(fields...)
	Logger.Debug(msg, fields...)
}

func (log Log) Info(msg string, fields ...zap.Field) {
	fields = log.AppendFields(fields...)
	Logger.Info(msg, fields...)
}

func (log Log) Warn(msg string, fields ...zap.Field) {
	fields = log.AppendFields(fields...)
	Logger.Warn(msg, fields...)
}

func (log Log) Error(msg string, fields ...zap.Field) {
	fields = log.AppendFields(fields...)
	Logger.Error(msg, fields...)
}

//在log中植入请求数据
func (log Log) AppendFields(fields ...zap.Field) []zap.Field {
	fields = append(fields, zap.String("trace_id", log.TraceId))
	return fields
}
