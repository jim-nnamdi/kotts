package user

import (
	"github.com/jim-nnamdi/kotts/internal/database"
	"go.uber.org/zap"
)

type user struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Country         string `json:"country"`
	Active          bool   `json:"active"`
	DatabaseHandler database.Client
	logger          *zap.Logger
}
