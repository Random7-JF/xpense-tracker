package db

import (
	"log"

	"github.com/Random7-JF/xpense-tracker/model"
)

func (s *Sqlite) FindExpenses(query string, parameters ...interface{}) ([]model.Expense, error) {
	result, err := s.Db.Query(query, parameters...)
	if err != nil {
		log.Printf("Error in Query for FindExpenses: %s", err)
		return []model.Expense{}, err
	}

	var expenses []model.Expense
	for result.Next() {
		var expense model.Expense
		err := result.Scan(&expense.Id, &expense.Label, &expense.Amount, &expense.Frequency, &expense.Tag,
			&expense.ExpenseDate, &expense.SubmissionDate, &expense.UserId)
		if err != nil {
			log.Printf("Error in RowScan of FindExpenses: %s - Current: %v", err, expense)
			return []model.Expense{}, err
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

func (s *Sqlite) GetExpense(userId string) ([]model.Expense, error) {
	query, err := ReadSQL("expense/getExpenses.sql")
	if err != nil {
		log.Println("Error in reading sql get expense", err)
		return nil, err
	}
	expenses, err := s.FindExpenses(query, userId)
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}
	return expenses, nil
}

func (s *Sqlite) GetExpenseByFreq(freq string, userid string) ([]model.Expense, error) {
	query, err := ReadSQL("expense/getExpenseByFreq.sql")
	if err != nil {
		log.Printf("Error reading the SQL for GetExpenseByFreq: %s", err)
		return []model.Expense{}, err
	}

	expenses, err := s.FindExpenses(query, freq, userid)
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}
	return expenses, nil
}

func (s *Sqlite) GetExpenseBySearch(search string, userid string) ([]model.Expense, error) {
	query, err := ReadSQL("expense/getExpenseBySearch.sql")
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseBySearch")
		return []model.Expense{}, err
	}
	expenses, err := s.FindExpenses(query, search, userid)
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}

	return expenses, nil
}

func (s *Sqlite) GetExpenseByTag(tag string, userid string) ([]model.Expense, error) {
	query, err := ReadSQL("expense/getExpensesByTag.sql")
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}
	expenses, err := s.FindExpenses(query, tag, userid)
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}

	return expenses, nil
}
