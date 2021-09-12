package controllers

import (
	"net/http"
	"new-bell/models"
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

	c.JSON(http.StatusOK, gin.H{"data": id})
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
	c.JSON(http.StatusOK, gin.H{"data": *user})
}

func GetUserInfoHandler(c *gin.Context) {
	uid := c.Param("id")
	user, err := service.GetUserInfo(uid)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": *user})
}
