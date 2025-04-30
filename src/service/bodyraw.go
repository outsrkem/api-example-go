package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outsrkem/api-example-go/src/logger"
)

type ReqRaw struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func TBodyRaw() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		klog := logger.FromContext(ctx)
		obj := ReqRaw{}
		if err := ctx.ShouldBindJSON(&obj); err != nil {
			klog.Errorf("%s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
			return
		}

		klog.Info("Bind ok.")
		klog.Debugf("%+v", obj)
		body := make(map[string]interface{})
		body["sss"] = obj
		ctx.JSON(http.StatusOK, body)
	}
}
