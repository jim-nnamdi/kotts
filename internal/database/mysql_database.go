package database

import (
	"database/sql"
	"errors"

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
	// config, err := loadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	var (
		err error
	)
	db, err = sql.Open("mysql", "root:M@etroboomin50@tcp(localhost:3306)/kotts")
	if err != nil {
		handler.logger.Debug("could not connect to the database")
		return
	}
	return db
}

func (handler *databaseHandler) GetUserByUsername(username string) bool {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where username = ?", username)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return false
	}
	return true
}

func (handler *databaseHandler) GetUserByEmail(email string) (*models.User, error) {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where email = ?", email)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return user_response, errors.New(err.Error())
	}
	return user_response, nil
}

func (handler *databaseHandler) GetByUsernameAndPassword(email string, password string) (*models.User, error) {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where email = ? and password = ?", email, password)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return nil, err
	}
	return user_response, nil
}

func (handler *databaseHandler) GetUserHash(email string) []byte {
	var (
		user_response = &models.User{}
		err           error
	)
	run_getsingleuser_query := handler.Databaseconn().QueryRow("select * from users where email = ?", email)
	if err = run_getsingleuser_query.Scan(
		&user_response.ID,
		&user_response.Username,
		&user_response.Password,
		&user_response.Country,
		&user_response.Email,
		&user_response.Active,
	); err != nil {
		return nil
	}
	return []byte(user_response.Password)
}
func (handler *databaseHandler) Close() error {
	return nil
}
