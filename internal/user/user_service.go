package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func RegistrationService(w http.ResponseWriter, r *http.Request) {
	var (
		username = r.FormValue("username")
		email    = r.FormValue("email")
		password = r.FormValue("password")
		country  = r.FormValue("country")
		active   = r.FormValue("active")
		user     = User{}
	)
	active_user_value, _ := strconv.Atoi(active)
	parse_form_for_registration, err := user.UserRegistration(username, email, password, country, active_user_value)
	if !parse_form_for_registration {
		log.Printf("Registration unsuccessful : %s", err)
		return
	} else {
		w.Write([]byte(username + " Registration successful!"))
	}
}

func LoginService(w http.ResponseWriter, r *http.Request) {
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
		user     = User{}
	)
	parse_form_to_login_user, err := user.UserLogin(email, password)
	if err != nil {
		log.Printf("login unsuccessful: %s", err)
		return
	}
	json.NewEncoder(w).Encode(parse_form_to_login_user)
}
