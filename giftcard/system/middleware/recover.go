package middleware

import (
	"errors"
	"giftcard/library/logger"
	"giftcard/library/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

//全局错误拦截器,可以在业务中随便panic,但是不建议这么干
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if c.IsAborted() {
					c.Status(200)
				}

				result := response.Result{
					Code:    response.STATUS_SERVER_ERROR,
					TraceId: c.GetString(logger.TraceId),
					Detail:  "此消息由Recover捕获",
				}

				switch errType := err.(type) {
				case string:
					result.Message = errType
				case error:
					wraperr := err.(error)
					errmsg := []string{}
					//最大10层
					for i := 0; i < 10; i++ {
						wraperr = errors.Unwrap(wraperr)
						if wraperr == nil {
							break
						}
						errmsg = append(errmsg, wraperr.Error())
					}

					result.Detail += strings.Join(errmsg, ",")
					result.Message = err.(error).Error()
				default:
					result.Message = "Unkonw Error"
				}

				logger.Error("请求错误", zap.String(logger.TraceId, c.GetString(logger.TraceId)), zap.Any("Error", result.Message))
				c.JSON(http.StatusOK, result)
			}
		}()
		c.Next()
	}
}
