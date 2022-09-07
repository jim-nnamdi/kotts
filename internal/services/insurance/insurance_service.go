package insurance

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func NewMobileInsuranceService(w http.ResponseWriter, r *http.Request) {
	var (
		name         = r.FormValue("name")
		email        = r.FormValue("email")
		phonenumber  = r.FormValue("phonenumber")
		nameofphone  = r.FormValue("nameofphone")
		purchasedate = r.FormValue("purchasedate")
		imeinumber   = r.FormValue("imeinumber")
		model        = r.FormValue("model")
		color        = r.FormValue("color")
		description  = r.FormValue("description")
		insurance    = &insurance{}
	)
	add_new_mobile_insurance, err := insurance.NewMobileInsurance(name, email, phonenumber, nameofphone, purchasedate, imeinumber, model, color, description, false, time.Now(), time.Now())
	if err != nil {
		log.Print(err.Error())
		return
	}
	if add_new_mobile_insurance {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nameofphone + " insurance data added successful!")
	}
}

func NewLaptopInsuranceService(w http.ResponseWriter, r *http.Request) {
	var (
		name         = r.FormValue("name")
		email        = r.FormValue("email")
		phonenumber  = r.FormValue("phonenumber")
		nameoflaptop = r.FormValue("nameoflaptop")
		purchasedate = r.FormValue("purchasedate")
		imeinumber   = r.FormValue("imeinumber")
		model        = r.FormValue("model")
		color        = r.FormValue("color")
		description  = r.FormValue("description")
		insurance    = &insurance{}
	)
	add_new_laptop_insurance, err := insurance.NewLaptopInsurance(name, email, phonenumber, nameoflaptop, purchasedate, imeinumber, model, color, description, false, time.Now(), time.Now())
	if err != nil {
		log.Print(err.Error())
		return
	}
	if add_new_laptop_insurance {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nameoflaptop + " insurance data added successful!")
	}
}
