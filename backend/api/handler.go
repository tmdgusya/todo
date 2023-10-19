package api

import (
	"context"
	"database/sql"
	"net/http"
)

// Handler is interface for http handler that go with database
type Hanlder interface {
	WebHanlder(ctx context.Context, db *sql.DB) func(w http.ResponseWriter, r *http.Request)
}
