package user

import "go.uber.org/zap"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Country  string `json:"country"`
	Active   int    `json:"active"`
	Logger   *zap.Logger
}

type Userinterface interface {
	UserRegistration(username string, email string, password string, country string, active int) (bool, error)
	UserLogin(email string, password string) (*User, error)
}
