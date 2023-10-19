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
