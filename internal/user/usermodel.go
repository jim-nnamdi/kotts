package user

type Userinterface interface {
	UserRegistration(username string, email string, password string, country string, active int) (bool, error)
}
