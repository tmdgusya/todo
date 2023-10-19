package repository

import (
	"database/sql"

	"tmdgusya.com/todo/model"
)

// InsertPost inserts a post into the database
func InsertPost(db *sql.DB, p *model.Post) int64 {
	stmt, err := db.Prepare("INSERT INTO posts(title, content, isChecked, createdAt, updatedAt, tagId, categoryId) VALUES(?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	res, err := stmt.Exec(p.Title, p.Content, p.IsChecked, p.CreatedAt, p.UpdatedAt, p.TagId, p.CategoryId)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	return lastId
}

func GetPosts(db *sql.DB, offset int) []model.Post {
	stmt, err := db.Prepare("SELECT * FROM posts limit 10 offset ?")
	var posts []model.Post = []model.Post{}
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.Query(offset)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.IsChecked, &post.CreatedAt, &post.UpdatedAt, &post.TagId, &post.CategoryId)

		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	return posts
}
