package router

import (
	"github.com/gin-gonic/gin"
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
}
