package main

import (
	"time"

	"github.com/edgex-go-api/caller"
	"github.com/edgex-go-api/config"
	"github.com/edgex-go-api/logs"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	config.InitConfig()
	logs.InitLogs()
	caller.InitClient()
}

func main() {

	gin.SetMode(config.ServerSetting.RunMode)

	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	registerRouter(r)

	_ = r.Run(config.ServerSetting.Port)
}
