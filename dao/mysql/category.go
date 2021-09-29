package mysql

import "new-bell/models"

func GetCategoryList(params *models.PageListParams, res *[]models.Category)( int,error)  {
	offset := (params.PageIndex -1) * params.PageSize
	err := db.Select(res, "SELECT cid,c_name FROM category LIMIT ? OFFSET ? ", params.PageSize, offset)
	if err != nil {
		return 0, err
	}
	count := []models.Total{}
	err = db.Select(&count, "SELECT COUNT(cid) as count FROM category")
	if err != nil {
		return 0, err
	}
	total := count[0].Count
	return total, nil
}