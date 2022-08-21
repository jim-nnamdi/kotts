package user

import (
	"errors"
	"fmt"
	"log"

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

var _ Userinterface = &User{}

func NewUser(username string, password string, email string, country string, active int) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
		Country:  country,
		Active:   active,
	}
}

func GenerateFromPassword(password string, cost int) ([]byte, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), cost)
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

func (usermodel *User) UserRegistration(username string, email string, password string, country string, active int) (bool, error) {
	if conn == nil {
		fmt.Println("connection could not be established")
	}
	res, err := conn.Prepare("insert into users (username,password, country, email, active) values(?,?,?,?,?)")
	if err != nil {
		// usermodel.Logger.Debug("could not run add user query successfully", zap.Error(err))
		return false, errors.New(err.Error())
	}
	hash_user_password, _ := GenerateFromPassword(password, 14)
	check_if_username_exists := db.GetUserByUsername(username)
	check_if_email_exists, err := db.GetUserByEmail(email)
	if err != nil {
		log.Print(err)
		return false, errors.New(err.Error())
	}
	if check_if_username_exists {
		return false, errors.New("user with username already exists")
	}
	if check_if_email_exists.ID != 0 {
		return false, errors.New("user with email already exists")
	}
	result_from_user_registration, err := res.Exec(username, hash_user_password, country, email, active)
	if err != nil {
		// usermodel.Logger.Debug("could not execute and insert user into database" + err.Error())
		return false, err
	}
	check_last_data_inserted, err := result_from_user_registration.RowsAffected()
	if check_last_data_inserted == 0 || check_last_data_inserted < 0 {
		return false, err
	}
	return true, nil
}

func (usermodel *User) UserLogin(email string, password string) (bool, error) {
	get_user_hash := db.GetUserHash(email)
	compare_hash_password, _ := CompareAndHashPassword(password, get_user_hash)
	if compare_hash_password {
		validate_user_login, err := db.GetByUsernameAndPassword(email, string(get_user_hash))
		if validate_user_login == nil {
			log.Print("userhandler : error logging in", err.Error())
			return false, errors.New("login failed: could not get user details")
		}
		return true, nil
	}
	return false, errors.New("error logging user in")
}
