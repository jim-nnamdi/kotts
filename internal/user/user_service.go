package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	jwt_secret_key = []byte("kotts_secret_key")
	user           = User{}
)

func RegistrationService(w http.ResponseWriter, r *http.Request) {
	var (
		username = r.FormValue("username")
		email    = r.FormValue("email")
		password = r.FormValue("password")
		country  = r.FormValue("country")
		active   = r.FormValue("active")
	)
	active_user_value, _ := strconv.Atoi(active)
	parse_form_for_registration, err := user.UserRegistration(username, email, password, country, active_user_value)
	if !parse_form_for_registration {
		log.Printf("Registration unsuccessful : %s", err)
		return
	} else {
		json.NewEncoder(w).Encode(username + " Registration successful!")
	}
}

func LoginService(w http.ResponseWriter, r *http.Request) {
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
	)
	parse_form_to_login_user, err := user.UserLogin(email, password)
	log.Print(parse_form_to_login_user)
	if !parse_form_to_login_user {
		log.Print("userservice : error logging in")
		return
	}
	if err != nil {
		log.Printf("login unsuccessful: %s", err)
		return
	}
	expiration_date := time.Now().Add(5 * time.Minute)
	parse_encoding_data := DataToEncode{
		Password: password,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_date.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, parse_encoding_data)
	token_string, err := token.SignedString(jwt_secret_key)
	if err != nil {
		log.Print(err.Error())
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "user_token",
		Value:   token_string,
		Expires: expiration_date,
	})
	json.NewEncoder(w).Encode(token_string)
}
