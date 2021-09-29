package service

import (
	"new-bell/dao/mysql"
	"new-bell/models"
)

func GetCategoryList(p *models.PageListParams) (*[]models.Category ,error) {
	clist := []models.Category{}
	err:= mysql.GetCategoryList(p, &clist)
	if err!=nil {
		return nil, err
	}
	return &clist, nil
}