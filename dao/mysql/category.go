package mysql

import "new-bell/models"

func GetCategoryList(params *models.PageListParams, res *[]models.Category)( err error)  {
	err = db.Select(res, "SELECT cid,c_name FROM category LIMIT ? OFFSET ? ", params.PageSize, params.PageIndex)
	if err != nil {
		return err
	}
	return
}
