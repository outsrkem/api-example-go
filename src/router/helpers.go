package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsrkem/api-example-go/src/logger"
)

func NoResponse(ctx *gin.Context) {
	klog := logger.FromContext(ctx)
	klog.Warnf("request a nonexistent api [%s %s]", ctx.Request.Method, ctx.Request.URL.Path)
	ctx.JSON(404, gin.H{
		"message": "api not exists!",
		"url":     ctx.Request.URL.Path,
		"method":  ctx.Request.Method,
	})
}

func helloWorld() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		klog := logger.FromContext(ctx)
		klog.Info("fun hello World！！")
		klog.Warn("Warn-----------")
		ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
	}
}
