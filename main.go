package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jim-nnamdi/kotts/internal/middlewares"
	"github.com/jim-nnamdi/kotts/internal/user"
)

func userreg(w http.ResponseWriter, r *http.Request) {
	user.RegistrationService(w, r)
}
func userlogin(w http.ResponseWriter, r *http.Request) {
	user.LoginService(w, r)
}

func testdocker(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("wonderful docker fr...")
}
func main() {
	log.Print("server started running at port 8080 ....")
	r := http.NewServeMux()
	testHandler := http.HandlerFunc(testdocker)
	userloginHandler := http.HandlerFunc(userlogin)
	userRegisterHandler := http.HandlerFunc(userreg)
	r.Handle("/test", middlewares.BearerMiddleware(testHandler))
	r.Handle("/login", userloginHandler)
	r.Handle("/register", userRegisterHandler)
	log.Print("listening on port 8080 ...")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
