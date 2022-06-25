package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

// Db is an wraping to sql.DB to create test mocks
type Db interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// GetConnection return connection from sql.DB
func GetConnection(context context.Context) *sql.DB {
	db, err := sql.Open("sqlite3", "./gymondo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}
