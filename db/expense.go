package db

import (
	"log"

	"github.com/Random7-JF/xpense-tracker/model"
)

func (s *Sqlite) GetExpense(userId string) ([]model.Expense, error) {
	query, err := ReadSQL("expense/getExpenses.sql")
	if err != nil {
		log.Println("Error in reading sql get expense", err)
		return nil, err
	}
	results, err := s.Db.Query(query, userId)
	if err != nil {
		log.Println("Error in query get expense", err)
		return nil, err
	}

	var expense model.Expense
	var expenses []model.Expense
	for results.Next() {
		err := results.Scan(&expense.Id, &expense.Label, &expense.Amount, &expense.Frequency, &expense.Tag,
			&expense.ExpenseDate, &expense.SubmissionDate, &expense.UserId)
		if err != nil {
			log.Println("Error in row scan for expense", err)
			return nil, err

		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (s *Sqlite) GetExpenseByID(expenseid string) (model.Expense, error) {
	query, err := ReadSQL("expense/getExpenseById.sql")
	if err != nil {
		log.Println("Error in reading sql get expense", err)
		return model.Expense{}, err
	}

	result := s.Db.QueryRow(query, expenseid)
	if err != nil {
		log.Println("QueryRow issue", err)
		return model.Expense{}, err
	}
	var expense model.Expense
	err = result.Scan(&expense.Id, &expense.Label, &expense.Amount, &expense.Frequency, &expense.Tag,
		&expense.ExpenseDate, &expense.SubmissionDate, &expense.UserId)
	if err != nil {
		log.Println("Scan Issue", err)
		return model.Expense{}, err
	}
	return expense, nil
}

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

func (s *Sqlite) GetExpenseByFreq(freq string) ([]model.Expense, error) {
	query, err := ReadSQL("expense/getExpenseByFreq.sql")
	if err != nil {
		log.Printf("Error reading the SQL for GetExpenseByFreq: %s", err)
		return []model.Expense{}, err
	}

	result, err := s.Db.Query(query, freq)
	if err != nil {
		log.Printf("Error in Query for GetExpenseByFreq: %s", err)
		return []model.Expense{}, err
	}

	var expenses []model.Expense
	for result.Next() {
		var expense model.Expense
		err := result.Scan(&expense.Id, &expense.Label, &expense.Amount, &expense.Frequency, &expense.Tag,
			&expense.ExpenseDate, &expense.SubmissionDate, &expense.UserId)
		if err != nil {
			log.Printf("Error in RowScan of GetExpenseByFreq: %s", err)
			return []model.Expense{}, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}
