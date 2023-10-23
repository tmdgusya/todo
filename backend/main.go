package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"regexp"

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
	server := api.RegexpServeMux{}
	server.AddRoute("GET", regexp.MustCompile("/api/health"), api.HandleApiHealthCheck)
	server.AddRoute("GET", regexp.MustCompile("/api/posts"), api.HandlePosts(context.Background(), db))
	server.AddRoute("GET", regexp.MustCompile("/api/categories"), api.HandleCategories(context.Background(), db))
	server.AddRoute("POST", regexp.MustCompile("/api/tags"), api.HandleTags(context.Background(), db))
	server.AddParamRoute("GET", regexp.MustCompile("/api/tags/([0-9]+)"), api.HandleGetTagById(context.Background(), db))
	http.Handle("/", &server)
	fmt.Printf("%+v\n", db)
	flag.Parse()

	http.ListenAndServe(*listenAddr, nil)

	defer db.Close()
}
