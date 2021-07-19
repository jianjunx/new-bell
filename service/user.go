package service

import (
	"errors"
	"fmt"
	"new-bell/dao/mysql"
	"new-bell/models"
)

func Signup(p *models.SignupParams) (int64, error) {
	var count []models.UserCount
	err := mysql.GetUserNameRowCount(p.UserName, &count)
	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}
	if count[0].UserCount != 0 {
		t := fmt.Sprintf(`用户名 "%s" 重复，请重新输入`, p.UserName)
		return 0, errors.New(t)
	}
	id, err := mysql.InsertUser(p)
	if err != nil {
		return 0, err
	}
	return id, nil
}
