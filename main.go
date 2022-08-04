package main

import (
	"fmt"
	"log"

	"github.com/jim-nnamdi/kotts/internal/user"
)

func main() {
	newuser := user.User{}
	x, err := newuser.UserRegistration("jim", "j@gmail.com", "123", "nga", 0)
	if !x {
		log.Print(err)
	}
	fmt.Print("User registered successfully ...")
}
