package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"tmdgusya.com/todo/db/repository"
	"tmdgusya.com/todo/model"
)

type CreateTagRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func HandleTags(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetTags(ctx, db)(w, r)
		case http.MethodPost:
			handleCreateTag(ctx, db)(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func HandleGetTagById(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request, matches []string) {
	return func(w http.ResponseWriter, r *http.Request, matches []string) {
		id := matches[1]
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(fmt.Errorf("id(%s) is not integer", id))
		}
		tag := repository.GetTagById(db, ctx, idInt)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(tag)))
	}
}

func handleGetTags(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleCreateTag(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := new(CreateTagRequest)
		json.NewDecoder(r.Body).Decode(request)
		id := repository.InsertTag(db, ctx, &model.Tag{Name: request.Name, Color: request.Color})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(id)))
	}
}
