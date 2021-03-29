package main

import (
	"net/http"
	"time"

	"github.com/edgex-go-api/logs"
	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		logs.Info("[Test] name=%v", "Bob")
		logs.Debug("[Test] name=%v", "Tim")
		c.JSON(http.StatusOK, gin.H{
			"msg":  "pong",
			"time": time.Now().Format(time.RFC3339),
		})
	})

}
