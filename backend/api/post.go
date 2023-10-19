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

type CreatePostRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	TagId      int64  `json:"tagId"`
	CategoryId int64  `json:"categoryId"`
}

func HandlePosts(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleCreatePost(ctx, db)(w, r)
		case http.MethodGet:
			handleGetPosts(ctx, db)(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func handleCreatePost(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := new(CreatePostRequest)

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			panic(err.Error())
		}
		fmt.Printf("Request : %+v\n", request)

		post := model.CreatePost(
			request.Title,
			request.Content,
			request.TagId,
			request.CategoryId,
		)

		id := repository.InsertPost(db, ctx, post)

		fmt.Printf("Post : %+v\n", post)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(id)))
	}
}

func handleGetPosts(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// need to receive offset from request
		posts := repository.GetPosts(db, ctx, 0)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(posts)))
	}
}
