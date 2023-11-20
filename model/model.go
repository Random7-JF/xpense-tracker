package model

import "database/sql"

type Expense struct {
	Id             int
	Label          string
	Amount         float64
	Frequency      string
	Tag            string
	ExpenseDate    string
	SubmissionDate string
	UserId         string
	Error          string
}

type User struct {
	Id            int
	Username      string
	Hasedpassword string
	Email         string
	Creationdate  string
	Lastlogin     sql.NullString
}
