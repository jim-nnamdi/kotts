package user

import (
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
