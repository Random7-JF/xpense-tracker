package db

import (
	"log"

	"github.com/Random7-JF/xpense-tracker/model"
)

func (s *Sqlite) AddExpense(e model.Expense) error {
	query, err := ReadSQL("expense/addExpense.sql")
	if err != nil {
		log.Println("Error in reading sql add expense:", err)
		return err
	}
	_, err = s.Db.Exec(query, e.Label, e.Amount, e.Frequency, e.Tag, e.ExpenseDate, e.SubmissionDate, e.UserId)
	if err != nil {
		log.Println("Addexpense error:", err)
		return err
	}
	return nil
}

func (s *Sqlite) RemoveExpense(expenseId int) error {
	query, err := ReadSQL("expense/removeExpense.sql")
	if err != nil {
		log.Println("Error in reading sql remove expense: ", err)
		return err
	}
	_, err = s.Db.Exec(query, expenseId)
	if err != nil {
		log.Println("Remove expense error: ", err)
	}

	return nil
}

func (s *Sqlite) UpdateExpenseById(expense model.Expense) error {
	query, err := ReadSQL("expense/updateExpenseById.sql")
	if err != nil {
		log.Printf("Error reading SQL for update expense by id:  %s", err)
		return err
	}
	_, err = s.Db.Exec(query, expense.Label, expense.Amount, expense.Frequency, expense.Tag, expense.Id)
	if err != nil {
		log.Printf("Error in query for update expense by id: %s", err)
		return err
	}

	return nil
}

func (s *Sqlite) ExpenseFill(userid string) error {
	query, err := ReadSQL("test/addExpenseToUserid.sql")
	if err != nil {
		log.Printf("Error reading the SQL for ExpenseFill: %s", err)
		return err
	}
	values := make([]interface{}, 26)
	for i := range values {
		values[i] = userid
	}
	_, err = s.Db.Exec(query, values...)
	if err != nil {
		log.Printf("Error executing sql for ExpenseFill %s", err)
		return err
	}

	return nil
}
