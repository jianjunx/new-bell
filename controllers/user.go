package controllers

import (
	"net/http"
	"new-bell/service"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	service.Signup()

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
