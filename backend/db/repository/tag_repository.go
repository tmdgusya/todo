package repository

import (
	"database/sql"

	"tmdgusya.com/todo/model"
)

// InsertTag inserts a tag into the database
func InsertTag(db *sql.DB, tag *model.Tag) int64 {
	stmt, err := db.Prepare("INSERT INTO tags(name, color) VALUES(?, ?)")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	res, err := stmt.Exec(tag.Name, tag.Color)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	return lastId
}

// GetTagById gets a tag by id
func GetTagById(db *sql.DB, id int64) model.Tag {
	stmt, err := db.Prepare("SELECT * FROM tags where id = ?")
	var tag model.Tag

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&tag.Id, &tag.Name, &tag.Color)

	if err != nil {
		panic(err.Error())
	}

	return tag
}
