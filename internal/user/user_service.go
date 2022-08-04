package user

import (
	"log"
	"net/http"
	"reflect"
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
	if reflect.DeepEqual(parse_form_for_registration, User{}) {
		log.Printf("Registration unsuccessful : %s", err)
		return
	}
	w.Write([]byte(err.Error()))
}
