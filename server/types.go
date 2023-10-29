package server

type RegisterForm struct {
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
	Error           string
}

type Auth struct {
	Valid   bool
	Message string
}
