package mysql

import (
	"database/sql"
	"new-bell/models"
	snowflake "new-bell/pkg/snowflak"
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
		listParams = append(listParams, val)
		listSql += " WHERE p.cid=?"
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
		countSql += " WHERE cid=?"
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

func GetPostsDetail(id string, posts *[]models.Posts) error {
	sql := `SELECT p.pid,p.title,p.content,p.cid,p.auth_id,p.create_at,p.update_at,c.c_name,u.user_name 
	FROM posts as p
	LEFT JOIN user as u ON p.auth_id=u.uid 
	LEFT JOIN category as c ON p.cid=c.cid
	WHERE p.pid=?`
	return db.Select(posts, sql, id)
}

func AddPost(param *models.AddPostParam, authId interface{}) (sql.Result, error) {
	pid := snowflake.GenerateId()
	sql := "INSERT INTO posts (pid,title,content,cid,auth_id) VALUES (?,?,?,?,?)"
	return db.Exec(sql, pid, param.Title, param.Content, param.Cid, authId)
}

func UpdatePost(param *models.AddPostParam, id interface{}) (sql.Result, error) {
	sql := "UPDATE posts title=?,content=?,cid=? WHERE pid=?"
	return db.Exec(sql, param.Title, param.Content, param.Cid, id)
}

func DeletePost(id interface{}) (sql.Result, error) {
	sql := "DELETE FROM posts WHERE pid=?"
	return db.Exec(sql, id)
}
