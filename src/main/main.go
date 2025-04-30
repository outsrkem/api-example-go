package main

import (
	"github.com/gin-gonic/gin"
	"github.com/outsrkem/api-example-go/src/logger"
	"github.com/outsrkem/api-example-go/src/router"
	"github.com/sirupsen/logrus"
)

func init() {
	logger.NewLogrus()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	svc := gin.Default()
	router.Index(svc)

	logrus.SetReportCaller(true)
	logrus.Info("xxxxxxxxxxxxx")
	logrus.Warn("Warn-----------")
	svc.Run()
}
