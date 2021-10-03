package routers

import (
	"new-bell/controllers"
	_ "new-bell/docs"
	"new-bell/logger"
	"new-bell/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 注册中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// swag path
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// 注册路由文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		// 注册
		v1.POST("/signup", controllers.SignupHandler)
		// 用户登录
		v1.POST("/login", controllers.LoginHandler)
		// 通过uid获取用户信息
		v1.GET("/user/:id", middleware.Identity(), controllers.GetUserInfoHandler)
		// 通过token获取用户信息
		v1.GET("/user/info", middleware.Identity(), controllers.GetUserDetail)

		// 获取分类列表
		v1.POST("/category", controllers.GetCategoryHandler)

		// 帖子 文章
		v1.POST("/posts/list", controllers.GetPostsHandler)
		// 获取文章详情
		v1.GET("/posts/:id", controllers.GetPostsDetailHandler)
		// 增加帖子
		v1.POST("/posts", middleware.Identity(), controllers.AddPostHandler)
		// 删除帖子
		v1.DELETE("/posts/:id", middleware.Identity(), controllers.DeletePostHandler)
	}

	return r
}
