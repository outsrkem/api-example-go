package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsrkem/api-example-go/src/logger"
)

func Tpaths() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		klog := logger.FromContext(ctx)
		name := ctx.Param("name")
		klog.Infof("%s", name)
		ctx.JSON(http.StatusOK, gin.H{"name": name})
	}
}
