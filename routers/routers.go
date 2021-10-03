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
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	v1 := r.Group("/v1")
	identity_v1 := v1.Group("", middleware.Identity())
	// 注册
	v1.POST("/signup", controllers.SignupHandler)
	// 用户登录
	v1.POST("/login", controllers.LoginHandler)
	// 通过uid获取用户信息
	identity_v1.GET("/user/:id", controllers.GetUserInfoHandler)
	// 通过token获取用户信息
	identity_v1.GET("/user/info", controllers.GetUserDetail)

	// 获取分类列表
	v1.POST("/category", controllers.GetCategoryHandler)

	// 帖子 文章
	v1.POST("/posts", controllers.GetPostsHandler)
	// 获取文章详情
	v1.GET("/posts/:id", controllers.GetPostsDetailHandler)
	return r
}
