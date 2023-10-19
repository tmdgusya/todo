package model

type Post struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsChecked  bool   `json:"isChecked"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	TagId      string `json:"tagId"`
	CategoryId int    `json:"categoryId"`
}
