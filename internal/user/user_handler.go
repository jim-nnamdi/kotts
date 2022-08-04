package user

import (
	"errors"

	"github.com/jim-nnamdi/kotts/internal/database"
	"go.uber.org/zap"
)

var (
	logger           *zap.Logger
	ErrAddingNewUser = errors.New("error adding new user to database")
	db               = database.NewDatabaseHandler(logger)
	conn             = db.Databaseconn()
)

func (usermodel *User) AddNew(username string, email string, password string, country string, active int) (bool, error) {
	res, err := conn.Prepare("insert into users (username,password, country, email, active) values(?,?,?,?,?)")
	if err != nil {
		usermodel.Logger.Debug("could not run add user query successfully", zap.Error(err))
		return false, errors.New(err.Error())
	}
	result_from_user_registration, err := res.Exec(username, password, country, email, active)
	if err != nil {
		usermodel.Logger.Debug("could not execute and insert user into database" + err.Error())
		return false, err
	}
	check_last_data_inserted, err := result_from_user_registration.RowsAffected()
	if check_last_data_inserted == 0 || check_last_data_inserted < 0 {
		return false, err
	}
	return true, nil
}
