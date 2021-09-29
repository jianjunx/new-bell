package controllers

import (
	"new-bell/models"
	"new-bell/service"

	"github.com/gin-gonic/gin"
)

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
