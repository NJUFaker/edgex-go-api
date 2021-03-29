package main

import (
	"net/http"
	"time"

	"github.com/edgex-go-api/logs"
	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.Engine) {

	logs.Info("[Test] name=%v", "Bob")
	logs.Debug("[Test] name=%v", "Tim")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "pong",
			"time": time.Now().Format(time.RFC3339),
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"status":  "",
			"message": "",
			"data":    []int64{1, 2, 3, 4, 5, 6},
		})
	})

}
