package database

import (
	"database/sql"
	"io"
	"log"

	"github.com/jim-nnamdi/kotts/internal/models"
	"github.com/jim-nnamdi/kotts/runner"
	"github.com/spf13/viper"
)

type Client interface {
	io.Closer
	Databaseconn() (db *sql.DB)
	GetUserByUsername(username string) bool
	GetUserByEmail(email string) (models.User, error)
	GetUserHash(email string) []byte
	GetByUsernameAndPassword(email string, password string) (*models.User, error)
}

func loadConfig(path string) (config runner.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Print(err.Error())
		return runner.Config{}, err
	}
	err = viper.Unmarshal(&config)
	return
}
