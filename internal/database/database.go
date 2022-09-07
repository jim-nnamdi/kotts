package database

import (
	"database/sql"
	"io"
	"time"

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

	// ----- mobile-phone insurance related methods ----- //
	ApplyForMobilePhoneInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error)

	// retrieve all insurance by user, filter by email
	AllMobilePhoneInsuranceApplications(email string) (*[]models.MobileInsurance, error)

	// single mobile insurance application
	SingleMobilePhoneInsurance(mobileinsuranceid int) (*models.MobileInsurance, error)

	// ----- laptop insurance related methods ---- //
	ApplyForLaptopInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error)

	// retrieve all insurance by user, filter by email
	AllLaptopsInsuranceApplications(email string) (*[]models.LaptopInsurance, error)

	// single mobile insurance application
	SingleLaptopInsurance(laptopinsuranceid int) (*models.LaptopInsurance, error)
}
