package mysql

import (
	"fmt"
	"new-bell/models"
)

func GetPosts(p *models.PageListParams, posts *[]models.Posts) (int, error) {
	listSql := `SELECT p.pid,p.title,p.content,p.cid,p.auth_id,p.create_at,p.update_at,c.c_name,u.user_name 
	FROM posts as p
	LEFT JOIN user as u ON p.auth_id=u.uid 
	LEFT JOIN category as c ON p.cid=c.cid`
	countSql := "SELECT COUNT(pid) as count FROM posts"

	listParams := []interface{}{}
	val, ok := p.Condition["cid"]
	if ok {
		where := fmt.Sprintf(" WHERE p.cid=%v", "?")
		listParams = append(listParams, val)
		listSql += where
		countSql += " WHERE cid=?"
	}
	listSql += " LIMIT ? OFFSET ?"
	offset := (p.PageIndex - 1) * p.PageSize
	listParams = append(listParams, p.PageSize, offset)
	err := db.Select(posts, listSql, listParams...)
	if err != nil {
		return 0, err
	}
	count := []models.Total{}
	if ok {
		err = db.Select(&count, countSql, val)
	} else {
		err = db.Select(&count, countSql)
	}
	if err != nil {
		return 0, err
	}
	total := count[0].Count
	return total, nil
}