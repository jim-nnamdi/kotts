package main

import (
	"fmt"

	"github.com/jim-nnamdi/kotts/internal/user"
)

func main() {
	newuser := &user.User{}
	new_user_info, _ := newuser.UserRegistration("jim", "jim@gmail.com", "123456", "Nigeria", 1)
	fmt.Print(new_user_info)
}
