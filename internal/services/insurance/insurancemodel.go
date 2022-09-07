package insurance

import (
	"time"

	"github.com/jim-nnamdi/kotts/internal/models"
)

type InsuranceInterface interface {
	NewMobileInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error)
	NewLaptopInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error)
	AllMobilePhoneInsuranceApplications(email string) (*[]models.MobileInsurance, error)
	AllLaptopsInsuranceApplications(email string) (*[]models.LaptopInsurance, error)
}

type MobileInsurance struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phonenumber  string    `json:"phonenumber"`
	Nameofphone  string    `json:"nameofphone"`
	Purchasedate string    `json:"purchasedate"`
	Imeinumber   string    `json:"imeinumber"`
	Model        string    `json:"model"`
	Color        string    `json:"color"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LaptopInsurance struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phonenumber  string    `json:"phonenumber"`
	Nameofphone  string    `json:"nameofphone"`
	Purchasedate string    `json:"purchasedate"`
	Imeinumber   string    `json:"imeinumber"`
	Model        string    `json:"model"`
	Color        string    `json:"color"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
