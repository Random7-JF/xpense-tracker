package db

import (
	"log"

	"github.com/Random7-JF/xpense-tracker/model"
)

func (s *Sqlite) GetExpense() ([]model.Expense, error) {
	query, err := ReadSQL("expense/getExpenses.sql")
	if err != nil {
		log.Println("Error in reading sql get expense", err)
		return nil, err
	}
	results, err := s.Db.Query(query)
	if err != nil {
		log.Println("Error in query get expense", err)
		return nil, err
	}

	var expense model.Expense
	var expenses []model.Expense
	for results.Next() {
		err := results.Scan(&expense.Id, &expense.Label, &expense.Amount, &expense.Tag,
			&expense.ExpenseDate, &expense.SubmissionDate, &expense.UserId)
		if err != nil {
			log.Println("Error in row scan for expense", err)
			return nil, err

		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (s *Sqlite) AddExpense(label string, amount float64, tags string, expenseDate string, submissionDate string, userId string) {
	query, err := ReadSQL("expense/addExpense.sql")
	if err != nil {
		log.Println("Error in reading sql add expense:", err)
		return
	}
	_, err = s.Db.Exec(query, label, amount, tags, expenseDate, submissionDate, userId)
	if err != nil {
		log.Println("Addexpense error:", err)
	}
}

func (s *Sqlite) RemoveExpense() {

}
