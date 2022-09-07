package database

import (
	"database/sql"
	"io"

	"github.com/jim-nnamdi/kotts/internal/models"
)

type Client interface {
	io.Closer
	Databaseconn() (db *sql.DB)
	GetUserByUsername(username string) bool
	GetUserByEmail(email string) (*models.User, error)
	GetUserHash(email string) []byte
	GetByUsernameAndPassword(email string, password string) (*models.User, error)

	// articles related methods

	GetAllArticles() (*[]models.Articles, error)
	GetByAuthor(author string) (*[]models.Articles, error)
	GetSingleArticle(articleID int) (*models.Articles, error)

	// mobile-phone insurance related methods
	ApplyForMobilePhoneInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string) (bool, error)
	AllMobilePhoneInsuranceApplications() (*[]models.MobileInsurance, error)
}
