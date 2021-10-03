package controllers

import (
	"new-bell/models"
	"new-bell/service"

	"github.com/gin-gonic/gin"
)

func GetPostsHandler(c *gin.Context) {
	p := models.PageListParams{}
	err := c.BindJSON(&p)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	posts, total, perr := service.GetPostsList(&p)
	if perr != nil {
		service.ErrorHandler(c, perr)
		return
	}
	service.SuccessList(c, &p, posts, total)
}
