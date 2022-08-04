package user

import (
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Country  string `json:"country"`
	Active   int    `json:"active"`
	Logger   *zap.Logger
}

type DataToEncode struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type Userinterface interface {
	UserRegistration(username string, email string, password string, country string, active int) (bool, error)
	UserLogin(email string, password string) (bool, error)
}
