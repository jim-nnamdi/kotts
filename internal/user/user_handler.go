package user

import (
	"errors"
	"log"

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
	x, er := res.Exec(username, password, country, email, active)
	if er != nil {
		log.Print(er)
		log.Print(x)
	}

	return true, nil
}
