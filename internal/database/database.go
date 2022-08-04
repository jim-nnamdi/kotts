package database

import (
	"database/sql"
	"io"

	"github.com/jim-nnamdi/kotts/internal/models"
)

type Client interface {
	io.Closer
	Databaseconn() (db *sql.DB)
	GetUserByUsername(username string) bool
	GetUserByEmail(email string) bool
	GetUserHash(email string) []byte
	GetByUsernameAndPassword(email string, password string) *models.User
}
