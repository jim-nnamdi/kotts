package user

import (
	"errors"

	"github.com/jim-nnamdi/kotts/internal/database"
	"go.uber.org/zap"
)

type User interface {
	AddNew(username string, email string, password string, country string, active bool) (bool, error)
}

var _ User = &user{}

var (
	ErrAddingNewUser = errors.New("error adding new user to database")
)

func NewUser(username string, email string, password string, country string, active bool, databaseClient database.Client, logger *zap.Logger) *user {
	return &user{
		Username:        username,
		Email:           email,
		Password:        password,
		Country:         country,
		Active:          active,
		DatabaseHandler: databaseClient,
		logger:          logger,
	}
}

func (handler *user) AddNew(username string, email string, password string, country string, active bool) (bool, error) {
	conn := handler.DatabaseHandler.Databaseconn()
	res, err := conn.Prepare("insert into users (username,password, active, country, email) values(?,?,?,?,?)")
	if err != nil {
		handler.logger.Debug("could not run add user query successfully", zap.Error(err))
		return false, errors.New(err.Error())
	}
	add_new_user_event, err := res.Exec(username, email, password, country, active)
	if err != nil {
		handler.logger.Debug("could not insert new user into the database")
		return false, errors.New(err.Error())
	}
	check_last_inserted_id, _ := add_new_user_event.LastInsertId()
	if check_last_inserted_id != 0 {
		return true, nil
	}
	return false, ErrAddingNewUser
}
