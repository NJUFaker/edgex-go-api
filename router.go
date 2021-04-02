package main

import (
	"net/http"
	"time"

	"github.com/edgex-go-api/caller"
	"github.com/edgex-go-api/handler"
	"github.com/edgex-go-api/logs"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
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

	r.GET("/test_redis", func(c *gin.Context) {
		ctx := context.Background()

		caller.RedisClient.Set(ctx, "test_redis", "success", 60*time.Second)
		val, err := caller.RedisClient.Get(ctx, "test_redis").Result()
		if err != nil && err != redis.Nil {
			logs.Error("[test_redis] redis get failed: key=%v, err=%v", "test_redis", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"status":  "",
			"message": "",
			"data": map[string]string{
				"key":   "test_redis",
				"value": val,
			},
		})
	})

	r.POST("/get_user_info", handler.GetUserInfo)

}
