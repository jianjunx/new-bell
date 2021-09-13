package middleware

import (
	"errors"
	"new-bell/pkg/jwt"
	"new-bell/service"
	"new-bell/settings"
	"strings"

	"github.com/gin-gonic/gin"
)

func Identity() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果在白名单中直接跳过
		if strings.Contains(settings.Config.WhiteList, c.FullPath()) {
			c.Next()
			return
		}
		// 取Header中的Token
		auth := c.GetHeader("Authorization")
		if auth == "" {
			service.ErrorHandler(c, errors.New(service.CodeNoPermission))
			return
		}
		token, claims, err := jwt.ParseToken(auth)
		if err != nil || !token.Valid {
			service.ErrorHandler(c, errors.New(service.CodeNoPermission))
			return
		}
		// 设置到上下文中
		c.Set("uid", claims.Uid)
		c.Set("userName", claims.UserName)
	}
}
