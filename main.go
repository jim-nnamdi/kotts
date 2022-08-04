package main

import (
	"log"
	"net/http"

	"github.com/jim-nnamdi/kotts/internal/user"
)

func userreg(w http.ResponseWriter, r *http.Request) {
	user.RegistrationService(w, r)
}
func userlogin(w http.ResponseWriter, r *http.Request) {
	user.LoginService(w, r)
}
func main() {
	log.Print("server started running at port 4500 ....")
	http.HandleFunc("/register", userreg)
	http.HandleFunc("/login", userlogin)
	log.Fatal(http.ListenAndServe(":4500", nil))
}
