package main

import (
	"flag"
	"net/http"

	"tmdgusya.com/todo/api"
)

func main() {
	listenAddr := flag.String("listen-addr", ":8080", "server listen address")
	flag.Parse()

	http.HandleFunc("/api/health", api.HandleApiHealthCheck)

	http.ListenAndServe(*listenAddr, nil)
}
