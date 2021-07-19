package mysql

import (
	"new-bell/models"
	snowflake "new-bell/pkg/snowflak"
)

func InsertUser(p *models.SignupParams) (int64, error) {
	uid := snowflake.GenerateId()
	_, err := db.Exec(
		"INSERT INTO user (uid, passwd, user_name) VALUES (?,?,?)",
		uid,
		p.Passwd,
		p.UserName)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func GetUserNameRowCount(username string, c *[]models.UserCount) (err error) {
	err = db.Select(c, "SELECT COUNT(user_name) as user_count FROM user WHERE user_name=?", username)
	return
}
