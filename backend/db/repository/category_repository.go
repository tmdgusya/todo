package repository

import (
	"database/sql"

	"tmdgusya.com/todo/model"
)

// InsertCategory inserts a category into the database
func InsertCategory(db *sql.DB, category *model.Category) int64 {
	stmt, err := db.Prepare("INSERT INTO categories(name) VALUES(?)")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	res, err := stmt.Exec(category.Name)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	return lastId
}
