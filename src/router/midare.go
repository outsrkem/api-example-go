package router

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/outsrkem/api-example-go/src/logger"
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
		klog := logger.FromContext(c)
		start := time.Now()
		c.Next()
		stop := time.Now()
		latency := stop.Sub(start)
		klog.Infof("|%14s | %d |%7s %s", latency, c.Writer.Status(), c.Request.Method, c.Request.RequestURI)
	}
}
