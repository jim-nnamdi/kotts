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
		w.Header().Set("Content-Type", "application/json")
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

	if err != nil {
		if !parse_form_to_login_user {
			log.Print("userservice : error logging in")
			return
		}
	} else {
		expiration_date := time.Now().Add(60 * time.Minute)
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

		// return user data
		user_data, err := db.GetUserByEmail(email)
		if err != nil {
			log.Print(err.Error())
			return
		}

		// data to be consumed
		// returned or by mobile app or me // sunday
		if user_data.BankDetails == nil && user_data.KYC == nil || user_data.BankDetails == nil || user_data.KYC == nil {
			user_result_data := map[string]interface{}{
				"personal_data": map[string]interface{}{
					"token":   token_string,
					"name":    user_data.Username,
					"email":   user_data.Email,
					"country": user_data.Country,
					"active":  user_data.Active,
				},
				"bank_details": map[string]interface{}{
					"account_name":   "",
					"account_number": "",
					"bank_name":      "",
					"bvn":            "",
				},
				"kyc": map[string]interface{}{
					"phone":   "",
					"address": "",
				},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user_result_data)
		} else {
			user_result_data := map[string]interface{}{
				"token":   token_string,
				"name":    user_data.Username,
				"email":   user_data.Email,
				"country": user_data.Country,
				"active":  user_data.Active,
				"bank_details": map[string]interface{}{
					"account_name":   user_data.BankDetails.AccountName,
					"account_number": user_data.BankDetails.AccountNumber,
					"bank_name":      user_data.BankDetails.BankName,
					"bvn":            user_data.BankDetails.BVN,
				},
				"kyc": map[string]interface{}{
					"phone":   user_data.KYC.Phone,
					"address": user_data.KYC.HomeAddress,
				},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user_result_data)
		}
	}
}
