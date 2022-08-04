package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jim-nnamdi/kotts/internal/models"
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

func (handler *databaseHandler) GetUserByUsername(ctx context.Context, username string) bool {
	var (
		user_response = &models.User{}
		err           error
	)
	get_user_by_username := fmt.Sprintf("select * from users where username = %s", username)
	run_getsingleuser_query := handler.Databaseconn().QueryRow(get_user_by_username)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		handler.logger.Debug("could not find user with username")
		return false
	}
	return true
}

func (handler *databaseHandler) Close() error {
	return nil
}
