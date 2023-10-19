package main

import (
	"flag"
	"fmt"
	"net/http"

	"tmdgusya.com/todo/api"
	"tmdgusya.com/todo/db"
)

func main() {
	listenAddr := flag.String("listen-addr", ":9000", "server listen address")
	db := db.Connection{
		Host:     "localhost",
		Port:     8881,
		User:     "root",
		Password: "roach",
		Database: "todos",
	}
	db.Connect()
	fmt.Printf("%+v\n", db)
	flag.Parse()

	http.HandleFunc("/api/health", api.HandleApiHealthCheck)

	http.ListenAndServe(*listenAddr, nil)
}
