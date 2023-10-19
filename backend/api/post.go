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

func HandleCreatePost(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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

		id := repository.InsertPost(db, post)

		fmt.Printf("Post : %+v\n", post)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(id)))
	}
}
