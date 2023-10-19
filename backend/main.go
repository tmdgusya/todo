package main

import (
	"flag"
	"fmt"
	"net/http"

	"tmdgusya.com/todo/api"
	"tmdgusya.com/todo/db"
)

func main() {
	listenAddr := flag.String("listen-addr", ":8080", "server listen address")
	db := db.Connection{
		Host:     "localhost",
		Port:     8801,
		User:     "root",
		Password: "roach",
		Database: "todo",
	}
	db.Connect()
	fmt.Printf("%+v\n", db)
	flag.Parse()

	http.HandleFunc("/api/health", api.HandleApiHealthCheck)

	http.ListenAndServe(*listenAddr, nil)
}
