package controllers

import (
	"fmt"
	"new-bell/models"
	"new-bell/pkg/jwt"
	"new-bell/service"

	"github.com/gin-gonic/gin"
)

// SignupHandler 注册接口
// @Summary 注册接口
// @Description 用户注册
// @Tags 注册
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body models.SignupParams false "参数"
// @Success 200 {string} 0
// @Router /signup [post]
func SignupHandler(c *gin.Context) {
	var p models.SignupParams
	err := c.ShouldBindJSON(&p)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	id, err := service.Signup(&p)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}

	service.SuccessData(c, id)
}

func LoginHandler(c *gin.Context) {
	var p models.SignupParams
	err := c.ShouldBindJSON(&p)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	user, err := service.Login(p.UserName, p.Passwd)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	token, err := jwt.Award(user, c.ClientIP())
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	service.SuccessData(c, &models.LoginResponse{Token: token, UserModal: *user})
}
func GetUserDetail(c *gin.Context) {
	uid, _ := c.Get("uid")
	user, err := service.GetUserInfo(fmt.Sprint(uid))
	if err != nil {
		service.ErrorHandler(c, err)
	}
	service.SuccessData(c, user)
}
func GetUserInfoHandler(c *gin.Context) {
	uid := c.Param("id")
	user, err := service.GetUserInfo(uid)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	service.SuccessData(c, user)
}
