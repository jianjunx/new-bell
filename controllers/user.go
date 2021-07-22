package controllers

import (
	"net/http"
	"new-bell/models"
	"new-bell/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignupHandler(c *gin.Context) {
	var p models.SignupParams
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("Params Err", zap.Error(err))
		return
	}
	id, err := service.Signup(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("Error", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

func GetUserInfoHandler(c *gin.Context) {
	uid := c.Query("uid")
	users, err := service.GetUserInfo(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("Error", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": *users})
}
