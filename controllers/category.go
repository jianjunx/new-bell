package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"new-bell/models"
	"new-bell/service"
)

func GetCategoryHandler(c *gin.Context)  {
	p:= models.PageListParams{
		PageSize: 10,
		PageIndex: 1,
	}
	err := c.BindJSON(p)
	if err != nil {
		return
	}
	list, err:= service.GetCategoryList(&p)
	if err!=nil {
		return
	}
	fmt.Print(list)
	//service.SuccessList(c, p, list,10)

}
