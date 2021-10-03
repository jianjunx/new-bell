package routers

import (
	"net/http"
	"new-bell/controllers"
	"new-bell/logger"
	"new-bell/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 注册中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Identity())
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	// 注册
	r.POST("/signup", controllers.SignupHandler)
	// 用户登录
	r.POST("/login", controllers.LoginHandler)
	// 通过uid获取用户信息
	r.GET("/user/:id", controllers.GetUserInfoHandler)
	// 通过token获取用户信息
	r.GET("/user/info", controllers.GetUserDetail)

	// 获取分类列表
	r.POST("/category", controllers.GetCategoryHandler)

	// 帖子 文章
	r.POST("/posts", controllers.GetPostsHandler)
	return r
}
