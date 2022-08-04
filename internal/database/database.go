package database

import (
	"database/sql"
	"io"
)

type Client interface {
	io.Closer
	Databaseconn() (db *sql.DB)
}
