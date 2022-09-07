package insurance

import (
	"errors"
	"time"

	"github.com/jim-nnamdi/kotts/internal/database"
	"go.uber.org/zap"
)

var _ InsuranceInterface = &insurance{}

type insurance struct {
	logger      *zap.Logger
	mysqlclient database.Client
}

func Newinsurance(logs *zap.Logger, db database.Client) *insurance {
	return &insurance{
		logger:      logs,
		mysqlclient: db,
	}
}

// new mobile insurance
func (ins *insurance) NewMobileInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error) {
	apply_for_new_mobile_insurance, err := ins.mysqlclient.ApplyForMobilePhoneInsurance(name, email, phonenumber, nameofphone, purchasedate, imeinumber, model, color, description, paid, createdAt, updatedAt)
	if err != nil {
		ins.logger.Debug("error applying for mobile insurance", zap.String("error", err.Error()))
		return false, err
	}
	if !apply_for_new_mobile_insurance {
		ins.logger.Debug("cannot add new mobile insurance application", zap.Bool("result", apply_for_new_mobile_insurance))
		return false, errors.New(err.Error())
	}
	return true, nil
}

// new laptop insurance
func (ins *insurance) NewLaptopInsurance(name string, email string, phonenumber string, nameofphone string, purchasedate string, imeinumber string, model string, color string, description string, paid bool, createdAt time.Time, updatedAt time.Time) (bool, error) {
	apply_for_new_laptop_insurance, err := ins.mysqlclient.ApplyForLaptopInsurance(name, email, phonenumber, nameofphone, purchasedate, imeinumber, model, color, description, paid, createdAt, updatedAt)
	if err != nil {
		ins.logger.Debug("error applying for laptop insurance", zap.String("error", err.Error()))
		return false, err
	}
	if !apply_for_new_laptop_insurance {
		ins.logger.Debug("cannot add new laptop insurance application", zap.Bool("result", apply_for_new_laptop_insurance))
		return false, errors.New(err.Error())
	}
	return false, nil
}
