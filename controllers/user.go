package controllers

import (
	"new-bell/models"
	"new-bell/pkg/jwt"
	"new-bell/service"

	"github.com/gin-gonic/gin"
)

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

func GetUserInfoHandler(c *gin.Context) {
	uid := c.Param("id")
	user, err := service.GetUserInfo(uid)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	service.SuccessData(c, user)
}
