package database

import (
	"context"
	"database/sql"
	"io"
)

type Client interface {
	io.Closer
	Databaseconn() (db *sql.DB)
	GetUserByUsername(ctx context.Context, username string) bool
}
