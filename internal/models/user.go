package models

import "go.uber.org/zap"

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Country     string `json:"country"`
	Active      int    `json:"active"`
	BankDetails *Bank  `json:"bank_details"`
	Logger      *zap.Logger
}

type Bank struct {
	ID            int    `json:"id"`
	AccountName   string `json:"account_name,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	BVN           string `json:"bvn,omitempty"`
	BankName      string `json:"bank_name,omitempty"`
}

type Userinterface interface {
	UserRegistration(username string, email string, password string, country string, active int) (bool, error)
}
