package main

import (
	"encoding/json"
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

func testdocker(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("wonderful docker fr...")
}
func main() {
	log.Print("server started running at port 8080 ....")
	r := mux.NewRouter()
	r.HandleFunc("/", testdocker)
	r.HandleFunc("/register", userreg)
	r.HandleFunc("/login", userlogin)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}
	log.Fatal(srv.ListenAndServe())
}
