package insurance

import "time"

var _ InsuranceInterface = &insurance{}

type insurance struct {
}

// new mobile insurance
func (ins *insurance) NewMobileInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error) {
	return false, nil
}

// new laptop insurance
func (ins *insurance) NewLaptopInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error) {
	return false, nil
}
