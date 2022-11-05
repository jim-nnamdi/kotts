package insurance

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		w.Header().Set("Accept", "application/json")
		json.NewEncoder(w).Encode(nameoflaptop + " insurance data added successful!")
	}
}

func AllMobileInsuranceApplicationService(w http.ResponseWriter, r *http.Request) {
	var (
		email     = r.FormValue("email")
		insurance = &insurance{}
	)

	return_users_mobileinsurance_applications, err := insurance.AllMobilePhoneInsuranceApplications(email)
	if err != nil {
		log.Print(err.Error())
		return
	}
	if return_users_mobileinsurance_applications != nil {
		log.Print("got here")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		json.NewEncoder(w).Encode(return_users_mobileinsurance_applications)
	}
}

func AllLaptopInsuranceApplicationService(w http.ResponseWriter, r *http.Request) {
	var (
		email     = r.FormValue("email")
		insurance = &insurance{}
	)

	return_users_laptopinsurance_applications, err := insurance.AllLaptopsInsuranceApplications(email)
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	json.NewEncoder(w).Encode(return_users_laptopinsurance_applications)
}

func SingleMobileInsuranceService(w http.ResponseWriter, r *http.Request) {
	var (
		insurance = &insurance{}
	)
	mob_insurance_id := r.URL.Query().Get("id")
	m_insurance_id, _ := strconv.Atoi(mob_insurance_id)
	single_insurance_data, err := insurance.SingleMobileInsurance(m_insurance_id)
	if err != nil {
		log.Print(err)
		return
	}
	json.NewEncoder(w).Encode(single_insurance_data)
}

func SingleLaptopInsuranceService(w http.ResponseWriter, r *http.Request) {
	var (
		insurance = &insurance{}
	)
	lap_insurance_id := r.URL.Query().Get("id")
	l_insurance_id, _ := strconv.Atoi(lap_insurance_id)
	single_insurance_data, err := insurance.SingleMobileInsurance(l_insurance_id)
	if err != nil {
		log.Print(err)
		return
	}
	json.NewEncoder(w).Encode(single_insurance_data)
}
