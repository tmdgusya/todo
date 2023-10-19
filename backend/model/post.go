package model

import "time"

type Post struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsChecked  bool   `json:"isChecked"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	TagId      int64  `json:"tagId"`
	CategoryId int64  `json:"categoryId"`
}

func CreatePost(
	title string,
	content string,
	tagId int64,
	categoryId int64,
) *Post {
	return &Post{
		Title:      title,
		Content:    content,
		IsChecked:  false,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		TagId:      tagId,
		CategoryId: categoryId,
	}
}
