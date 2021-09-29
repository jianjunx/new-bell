package service

import (
	"new-bell/dao/mysql"
	"new-bell/models"
)

func GetCategoryList(p *models.PageListParams) (*[]models.Category, int, error) {
	clist := []models.Category{}
	total, err := mysql.GetCategoryList(p, &clist)
	if err != nil {
		return nil, 0, err
	}
	return &clist, total, nil
}
