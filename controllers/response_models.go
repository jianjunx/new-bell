package controllers

import "new-bell/models"

type Res struct {
	Code int
	Data int
	Msg  string
}

type CategoryRes struct {
	Code      int
	Msg       string
	List      []models.Category
	Total     int
	PageSize  int
	PageIndex int
}

type PostsRes struct {
	Code      int
	Msg       string
	List      []models.Posts
	Total     int
	PageSize  int
	PageIndex int
}