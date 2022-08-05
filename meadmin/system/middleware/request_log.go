package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"meadmin/library/logger"
	"meadmin/library/strings"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		traceId := c.GetHeader(logger.TraceId)
		if len(traceId) <= 0 {
			traceId = c.GetHeader("RequestId")
		}

		if len(traceId) <= 0 {
			traceId = strings.GenUuid()
			c.Header(logger.TraceId, traceId)
			c.Set(logger.TraceId, traceId)
		}

		body, _ := c.GetRawData()
		logger.Info("Request:", zap.String("Method", c.Request.Method),
			zap.String("Path", c.FullPath()),
			zap.String("Body", string(body)),
			zap.String("TraceId", traceId))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		c.Next()
		responseBody := bodyLogWriter.body.String()
		logger.Info("Response:", zap.String("Body", responseBody), zap.String("TraceId", traceId))
	}
}
