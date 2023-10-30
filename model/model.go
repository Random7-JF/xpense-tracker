package model

type Expense struct {
	Id             int
	Label          string
	Amount         float64
	Tag            string
	ExpenseDate    string
	SubmissionDate string
	UserId         string
	Error          string
}
