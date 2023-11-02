package model

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
