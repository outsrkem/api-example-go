package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoResponse(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "api not exists!",
		"url":     c.Request.URL.Path,
		"method":  c.Request.Method,
	})
}

func helloWorld() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	}
}
