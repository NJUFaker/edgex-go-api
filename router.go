package main

import (
	"github.com/edgex-go-api/handler"
	"github.com/edgex-go-api/wrapper"
	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.Engine) {
	// TEST
	r.POST("/get_user_info", wrapper.JsonOutPutWrapper(handler.GetUserInfo))
	// your code

	appRouter := r.Group("/edgex/app")
	{
		appRouter.GET("/device", wrapper.JsonOutPutWrapper(handler.GetAllDevice))
	}
}
