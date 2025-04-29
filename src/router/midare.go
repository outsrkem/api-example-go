package router

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = strings.ReplaceAll(uuid.New().String(), "-", "")
		}
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}

func RequestRrecorder() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func CustomLogFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 从上下文中获取 RequestID
		requestID := "-"
		if val, exists := param.Keys["RequestID"]; exists {
			if id, ok := val.(string); ok {
				requestID = id
			}
		}
		// 自定义格式
		return fmt.Sprintf("[%s] [%s] \"%s %s %s %d %s\" \"%s\" \n",
			param.TimeStamp.Format(time.DateTime),
			requestID,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	})
}
