package router

import (
	"github.com/gin-gonic/gin"
	"github.com/outsrkem/api-example-go/src/service"
)

func midare(r *gin.Engine) {
	r.Use(RequestIDMiddleware())
	r.Use(RequestRrecorder())
	// r.Use(CustomLogFormat())
}

func Index(r *gin.Engine) {
	midare(r)
	r.NoRoute(NoResponse)
	r.GET("/", helloWorld())
	r.GET("/ping", helloWorld())
	r.POST("/body/raw", service.TBodyRaw())
	r.GET("/path/:name", service.Tpaths())
	r.GET("/query", service.Tquerp())
}
