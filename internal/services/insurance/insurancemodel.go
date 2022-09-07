package insurance

import "time"

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
