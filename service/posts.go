package service

import (
	"new-bell/dao/mysql"
	"new-bell/models"
)

func GetPostsList(p *models.PageListParams) (*[]models.Posts, int, error) {
	posts := []models.Posts{}
	total, err := mysql.GetPosts(p, &posts)
	if err != nil {
		return nil, 0, err
	}
	return &posts, total, nil
}

func GetPostsDetail(id string) (*models.Posts, error) {
	posts := []models.Posts{}
	err := mysql.GetPostsDetail(id, &posts)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, nil
	}
	return &posts[0], nil
}
