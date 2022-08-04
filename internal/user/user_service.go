package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
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

	type DataToEncode struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		jwt.StandardClaims
	}
	expiration_date := time.Now().Add(5 * time.Minute)
	jwt_secret_key := []byte("kotts_secret_key")
	parse_encoding_data := DataToEncode{
		Username: parse_form_to_login_user.Username,
		Email:    parse_form_to_login_user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_date.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, parse_encoding_data)
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
}
