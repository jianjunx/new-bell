package models

type UserModal struct {
	Uid      int    `db:"uid" json:"uid"`
	UserName string `db:"user_name" json:"userName"`
	Email    string `db:"email" json:"email"`
	Gender   string `db:"gender" json:"gender"`
	CreateAt string `db:"create_at" json:"createAt"`
	UpdateAt string `db:"update_at" json:"updateAt"`
}
type UserCount struct {
	UserCount int `db:"user_count"`
}
