package controllers

import (
	"new-bell/models"
	"new-bell/service"

	"github.com/gin-gonic/gin"
)

// GetCategoryHandler 获取分类列表
// @Summary 获取分类列表
// @Description 分类列表
// @Tags category
// @Accept application/json
// @Produce application/json
// @Param object body models.PageListParams true "参数"
// @Success 200 {object} CategoryRes
// @Router /category [post]
func GetCategoryHandler(c *gin.Context) {
	p := models.PageListParams{
		PageSize:  10,
		PageIndex: 1,
	}
	err := c.BindJSON(&p)
	if err != nil {
		service.ErrorHandler(c, err)
		return
	}
	list, total, cerr := service.GetCategoryList(&p)
	if cerr != nil {
		service.ErrorHandler(c, cerr)
		return
	}
	service.SuccessList(c, &p, &list, total)
}
