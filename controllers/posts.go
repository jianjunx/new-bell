package controllers

import (
	"errors"
	"new-bell/models"
	"new-bell/service"

	"github.com/gin-gonic/gin"
)

// GetPostsHandler 获取文章列表
// @Summary 获取文章列表
// @Description 文章列表
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param object body models.PageListParams true "参数"
// @Success 200 {object} PostsRes
// @Router /posts/list [post]
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

// GetPostsDetailHandler 获取文章详情
// @Summary 获取文章详情
// @Description 文章详情
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param id path int true "参数"
// @Success 200 {object} models.Posts
// @Router /posts/:id [get]
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

// AddPostHandler 添加文章
// @Summary 添加文章
// @Description 添加文章
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param object body models.AddPostParam true "参数"
// @Success 200 {object} Res
// @Router /posts [post]
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

// UpdatePostHandler 更新文章
// @Summary 更新文章
// @Description 更新文章
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param object body models.AddPostParam true "参数"
// @Param id path int true "pid"
// @Success 200 {object} Res
// @Router /posts/:id [put]
func UpdatePostHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		service.ErrorHandler(c, errors.New("ID不能为空"))
		return
	}
	param := models.AddPostParam{}
	err := c.BindJSON(&param)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	err = service.UpdatePost(&param, &id)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	service.SuccessData(c, nil)
}

// GetPostsDetailHandler 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param id path int true "参数"
// @Success 200 {object} Res
// @Router /posts/:id [delete]
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
