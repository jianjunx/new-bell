package models

type Posts struct {
	Pid      int    `db:"pid" json:"pid"`
	Title    string `db:"title" json:"title"`
	Content  string `db:"content" json:"content"`
	Cid      int    `db:"cid" json:"cid"`
	Cname    string `db:"c_name" json:"cname"`
	AuthID   int    `db:"auth_id" json:"authId"`
	AuthName string `db:"user_name" json:"authName"`
	CreateAt string `db:"create_at" json:"createAt"`
	UpdateAt string `db:"update_at" json:"updateAt"`
}
