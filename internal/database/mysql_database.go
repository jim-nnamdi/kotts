package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var _ Client = &databaseHandler{}

type databaseHandler struct {
	logger *zap.Logger
}

func NewDatabaseHandler(logger *zap.Logger) *databaseHandler {
	return &databaseHandler{
		logger: logger,
	}
}

func (handler *databaseHandler) Databaseconn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:M@etroboomin50@tcp(localhost:3306)/kotts")
	if err != nil {
		handler.logger.Debug("could not connect to the database")
		return
	}
	return db
}

func (handler *databaseHandler) Close() error {
	return nil
}
