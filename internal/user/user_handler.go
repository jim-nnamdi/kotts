package user

import (
	"errors"

	"github.com/jim-nnamdi/kotts/internal/database"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	logger           *zap.Logger
	ErrAddingNewUser = errors.New("error adding new user to database")
	db               = database.NewDatabaseHandler(logger)
	conn             = db.Databaseconn()
)

func GenerateFromPassword(password string, cost int) ([]byte, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	return hashed_password, nil
}

func CompareAndHashPassword(password string, hash []byte) (bool, error) {
	convert_hash_to_pwd := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if convert_hash_to_pwd != nil {
		return false, errors.New(convert_hash_to_pwd.Error())
	}
	return true, nil
}

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
