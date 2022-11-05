package main

import (
	"log"
	"net/http"

	"github.com/jim-nnamdi/kotts/internal/middlewares"
	"github.com/jim-nnamdi/kotts/internal/services/insurance"
	"github.com/jim-nnamdi/kotts/internal/services/user"
)

// serves as an entrypoint & display msg onload
func entrypoint(w http.ResponseWriter, r *http.Request) {
	log.Print("connection opened for entrypoint ...")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello kotts"))
}

func userreg(w http.ResponseWriter, r *http.Request) {
	user.RegistrationService(w, r)
}
func userlogin(w http.ResponseWriter, r *http.Request) {
	user.LoginService(w, r)
}

func newMobileInsuranceService(w http.ResponseWriter, r *http.Request) {
	insurance.NewMobileInsuranceService(w, r)
}

func newLaptopInsuranceService(w http.ResponseWriter, r *http.Request) {
	insurance.NewLaptopInsuranceService(w, r)
}

func allMobileInsuranceApplicationService(w http.ResponseWriter, r *http.Request) {
	insurance.AllMobileInsuranceApplicationService(w, r)
}

func allLaptopInsuranceApplicationService(w http.ResponseWriter, r *http.Request) {
	insurance.AllLaptopInsuranceApplicationService(w, r)
}

func singleMobileInsuranceService(w http.ResponseWriter, r *http.Request) {
	insurance.SingleMobileInsuranceService(w, r)
}

func singleLaptopInsuranceService(w http.ResponseWriter, r *http.Request) {
	insurance.SingleLaptopInsuranceService(w, r)
}

func main() {
	log.Print("server started running at port 8080 ....")
	r := http.NewServeMux()

	// entrypoint
	epserve := http.HandlerFunc(entrypoint)

	// users endpoints
	userloginHandler := http.HandlerFunc(userlogin)
	userRegisterHandler := http.HandlerFunc(userreg)

	// insurance endpoints
	newMobileInsuranceService := http.HandlerFunc(newMobileInsuranceService)
	newLaptopInsuranceService := http.HandlerFunc(newLaptopInsuranceService)
	allMobileInsuranceApplicationService := http.HandlerFunc(allMobileInsuranceApplicationService)
	allLaptopInsuranceApplicationService := http.HandlerFunc(allLaptopInsuranceApplicationService)
	singleMobileInsuranceService := http.HandlerFunc(singleMobileInsuranceService)
	singleLaptopInsuranceService := http.HandlerFunc(singleLaptopInsuranceService)

	// auth handlers

	r.Handle("/", middlewares.Jwtmiddleware(epserve))
	r.Handle("/login", middlewares.RecoveryMiddleware(userloginHandler))
	r.Handle("/register", middlewares.RecoveryMiddleware(userRegisterHandler))

	// insurance handlers

	r.Handle("/new-mobile-insurance", middlewares.RecoveryMiddleware(newMobileInsuranceService))

	r.Handle("/new-laptop-insurance", middlewares.RecoveryMiddleware(newLaptopInsuranceService))

	r.Handle("/all-mobile-insurance-applications", middlewares.RecoveryMiddleware(allMobileInsuranceApplicationService))

	r.Handle("/all-laptop-insurance-applications", middlewares.RecoveryMiddleware(allLaptopInsuranceApplicationService))

	r.Handle("/mobile-insurance", middlewares.RecoveryMiddleware(singleMobileInsuranceService))

	r.Handle("/laptop-insurance", middlewares.RecoveryMiddleware(singleLaptopInsuranceService))

	log.Print("listening on port 8080 ...")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
