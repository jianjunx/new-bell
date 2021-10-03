package controllers

import (
	"errors"
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

func GetPostsDetailHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		service.ErrorHandler(c, errors.New("ID不能为空"))
		return
	}
	post, err := service.GetPostsDetail(id)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	service.SuccessData(c, post)
}

func AddPostHandler(c *gin.Context) {
	param := models.AddPostParam{}
	authId, _ := c.Get("uid")
	err := c.BindJSON(&param)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	err = service.AddPost(&param, authId)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	service.SuccessData(c, nil)
}

func DeletePostHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		service.ErrorHandler(c, errors.New("ID不能为空"))
		return
	}
	err := service.DeletePost(&id)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	service.SuccessData(c, nil)
}
