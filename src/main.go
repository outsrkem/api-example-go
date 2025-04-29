package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoResponse(c *gin.Context) {
	// 返回 404 状态码
	c.String(http.StatusNotFound, "404, page not exists!")
}

func main() {

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.NoRoute(NoResponse)
	router.Run()
}
