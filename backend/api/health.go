package api

import "net/http"

func HandleApiHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
