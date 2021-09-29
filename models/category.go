package models

type Category struct {
	Cid int `json:"cid" db:"cid"`
	Cname string `json:"cname" db:"c_name"`
}
