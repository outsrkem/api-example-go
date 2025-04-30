package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsrkem/api-example-go/src/logger"
)

func Tquerp() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		klog := logger.FromContext(ctx)
		name := ctx.Query("name")
		age := ctx.Query("age")
		klog.Infof("name=%s age=%s", name, age)
		ctx.Header("add-header", "12314")
		ctx.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	}
}
