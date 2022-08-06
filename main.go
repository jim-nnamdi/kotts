package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jim-nnamdi/kotts/internal/user"
)

func userreg(w http.ResponseWriter, r *http.Request) {
	user.RegistrationService(w, r)
}
func userlogin(w http.ResponseWriter, r *http.Request) {
	user.LoginService(w, r)
}
func main() {
	log.Print("server started running at port 7800 ....")
	r := mux.NewRouter()
	r.HandleFunc("/register", userreg)
	r.HandleFunc("/login", userlogin)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:7800",
	}
	log.Fatal(srv.ListenAndServe())
}
