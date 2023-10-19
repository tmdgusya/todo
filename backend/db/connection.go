package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// Connect to MySQL database with given information
func (c *Connection) Connect() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
