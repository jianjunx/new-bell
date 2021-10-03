package models

type SignupParams struct {
	UserName string `form:"userName" binding:"required"`
	Passwd   string `form:"passwd" binding:"required"`
	Repasswd string `form:"repasswd"`
}

type PageListParams struct {
	PageIndex int                    `form:"pageIndex"`
	PageSize  int                    `form:"pageSize"`
	Condition map[string]interface{} `form:"condition"`
	Asc       []string               `form:"asc"`
	Desc      []string               `form:"desc"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	UserModal `json:"userInfo"`
}

type Total struct {
	Count int `json:"count"`
}

type AddPostParam struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Cid     int    `form:"cid" json:"cid" binding:"required"`
}
