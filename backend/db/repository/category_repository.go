package repository

import (
	"context"
	"database/sql"

	"tmdgusya.com/todo/model"
)

// InsertCategory inserts a category into the database
func InsertCategory(db *sql.DB, ctx context.Context, category *model.Category) int64 {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO categories(name) VALUES(?)")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, category.Name)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	return lastId
}

func GetCategories(db *sql.DB, ctx context.Context) []model.Category {
	stmt, err := db.PrepareContext(ctx, "SELECT * FROM categories")
	var categories []model.Category = []model.Category{}

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)

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
