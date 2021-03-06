package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanxiaox/ginx/internal/api"
	"github.com/xuanxiaox/ginx/internal/middleware"
)

func Router(r *gin.Engine) {
	r.Use(middleware.Cors())
	v1 := r.Group("/user")
	{
		user := api.NewUser()
		//v1.GET("/:id", user.Get)
		v1.GET("/list", user.List)
	}
}
