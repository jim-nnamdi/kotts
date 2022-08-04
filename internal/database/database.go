package database

import (
	"database/sql"
	"io"
)

type Client interface {
	io.Closer
	Databaseconn() (db *sql.DB)
	GetUserByUsername(username string) bool
	GetUserByEmail(email string) bool
	GetUserHash(email string) []byte
	GetByUsernameAndPassword(email string, password string) bool
}
