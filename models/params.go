package models

type SignupParams struct {
	UserName string `form:"user_name" binding:"required"`
	Passwd   string `form:"passwd" binding:"required"`
	Repasswd string `form:"repasswd"`
}
