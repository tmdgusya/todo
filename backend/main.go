package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"tmdgusya.com/todo/api"
	"tmdgusya.com/todo/db"
)

func main() {
	listenAddr := flag.String("listen-addr", ":9000", "server listen address")
	connection := db.Connection{
		Host:     "localhost",
		Port:     8881,
		User:     "root",
		Password: "roach",
		Database: "todos",
	}
	db := connection.Connect()
	fmt.Printf("%+v\n", db)
	flag.Parse()

	http.HandleFunc("/api/health", api.HandleApiHealthCheck)
	http.HandleFunc("/api/posts", api.HandlePosts(context.Background(), db))
	http.HandleFunc("/api/categories", api.HandleCategories(context.Background(), db))

	http.ListenAndServe(*listenAddr, nil)

	defer db.Close()
}
