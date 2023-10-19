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

func GetCategories(db *sql.DB) []model.Category {
	stmt, err := db.Prepare("SELECT * FROM categories")
	var categories []model.Category = []model.Category{}

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
	}

	return categories
}
