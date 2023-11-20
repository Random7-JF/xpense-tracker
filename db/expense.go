package db

import (
	"log"

	"github.com/Random7-JF/xpense-tracker/model"
)

func (s *Sqlite) AddExpense(e model.Expense) error {
	_, err := s.Db.Exec(s.Sql["sql/expense/addExpense.sql"], e.Label, e.Amount, e.Frequency, e.Tag, e.ExpenseDate, e.SubmissionDate, e.UserId)
	if err != nil {
		log.Println("Addexpense error:", err)
		return err
	}
	return nil
}

func (s *Sqlite) RemoveExpense(expenseId int) error {
	_, err := s.Db.Exec(s.Sql["sql/expense/removeExpense.sql"], expenseId)
	if err != nil {
		log.Println("Remove expense error: ", err)
	}

	return nil
}

func (s *Sqlite) UpdateExpenseById(expense model.Expense) error {
	_, err := s.Db.Exec(s.Sql["sql/expense/updateExpenseById.sql"], expense.Label, expense.Amount, expense.Frequency, expense.Tag, expense.Id)
	if err != nil {
		log.Printf("Error in query for update expense by id: %s", err)
		return err
	}

	return nil
}

func (s *Sqlite) ExpenseFill(userid string) error {
	values := make([]interface{}, 26)
	for i := range values {
		values[i] = userid
	}
	_, err := s.Db.Exec(s.Sql["sql/test/addExpenseToUserid.sql"], values...)
	if err != nil {
		log.Printf("Error executing sql for ExpenseFill %s", err)
		return err
	}

	return nil
}
