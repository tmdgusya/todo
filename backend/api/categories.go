package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"tmdgusya.com/todo/db/repository"
	"tmdgusya.com/todo/model"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func HandleCategories(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetCategories(ctx, db)(w, r)
		case http.MethodPost:
			handleCreateCategory(ctx, db)(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func handleGetCategories(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleCreateCategory(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := new(CreateCategoryRequest)
		json.NewDecoder(r.Body).Decode(request)
		id := repository.InsertCategory(db, &model.Category{Name: request.Name})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(id)))
	}
}
