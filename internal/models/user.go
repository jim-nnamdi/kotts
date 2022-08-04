package models

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
