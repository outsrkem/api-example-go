package main

import (
	"github.com/outsrkem/api-example-go/src/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	svc := gin.Default()
	router.Index(svc)
	svc.Run()
}
