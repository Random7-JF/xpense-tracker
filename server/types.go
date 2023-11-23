package server

type RegisterForm struct {
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
	Error           string
}

type LoginForm struct {
	Username string
	Password string
	Error    string
}

type Auth struct {
	Valid    bool
	Username string
	UserId   string
	Message  string
}
