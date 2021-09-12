package models

type SignupParams struct {
	UserName string `form:"userName" binding:"required"`
	Passwd   string `form:"passwd" binding:"required"`
	Repasswd string `form:"repasswd"`
}

type PageListParams struct {
	PageIndex int         `form:"pageIndex"`
	PageSize  int         `form:"pageSize"`
	Condition interface{} `form:"condition"`
	Asc       []string    `form:"asc"`
	Desc      []string    `form:"desc"`
}
