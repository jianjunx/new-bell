package routers

import (
	"net/http"
	"new-bell/controllers"
	"new-bell/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 注册中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	// 注册
	r.POST("/signup", controllers.SignupHandler)

	return r
}
