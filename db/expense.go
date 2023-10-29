package db

import (
	"log"
)

func (s *Sqlite) GetExpense() {

}

func (s *Sqlite) AddExpense(label string, amount float64, tags string, expenseDate string, submissionDate string, userId string) {
	query, err := ReadSQL("expense/addExpense.sql")
	if err != nil {
		log.Println("Error in reading sql add expense:", err)
		return
	}

	result, err := s.Db.Exec(query, label, amount, tags, expenseDate, submissionDate, userId)
	lines, _ := result.RowsAffected()
	log.Println("Add Expense:", lines)
}

func (s *Sqlite) RemoveExpense() {

}
