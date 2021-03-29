package main

import (
	"time"

	"github.com/edgex-go-api/config"
	"github.com/edgex-go-api/logs"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	config.InitAppConfig()
	logs.InitLogs()
}

func main() {

	gin.SetMode(config.AppConfig.RunMode)

	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	registerRouter(r)

	r.Run(config.AppConfig.Port)
}
