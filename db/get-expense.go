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
	result := s.Db.QueryRow(s.Sql["sql/expense/getExpenseById.sql"], expenseid)
	if result.Err() != nil {
		log.Println("QueryRow issue", result.Err())
		return model.Expense{}, result.Err()
	}
	var expense model.Expense
	err := result.Scan(&expense.Id, &expense.Label, &expense.Amount, &expense.Frequency, &expense.Tag,
		&expense.ExpenseDate, &expense.SubmissionDate, &expense.UserId)
	if err != nil {
		log.Println("Scan Issue", err)
		return model.Expense{}, err
	}
	return expense, nil
}

func (s *Sqlite) GetExpense(userId string) ([]model.Expense, error) {
	expenses, err := s.FindExpenses(s.Sql["sql/expense/getExpenses.sql"], userId)
	if err != nil {
		log.Printf("Error reading the SQL for getExpense")
		return []model.Expense{}, err
	}
	return expenses, nil
}

func (s *Sqlite) GetExpenseByFreq(freq string, userid string) ([]model.Expense, error) {
	expenses, err := s.FindExpenses(s.Sql["sql/expense/getExpenseByFreq.sql"], freq, userid)
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}
	return expenses, nil
}

func (s *Sqlite) GetExpenseBySearch(search string, userid string) ([]model.Expense, error) {
	expenses, err := s.FindExpenses(s.Sql["sql/expense/getExpenseBySearch.sql"], search, userid)
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}

	return expenses, nil
}

func (s *Sqlite) GetExpenseByTag(tag string, userid string) ([]model.Expense, error) {
	expenses, err := s.FindExpenses(s.Sql["sql/expense/getExpensesByTag.sql"], tag, userid)
	if err != nil {
		log.Printf("Error reading the SQL for getExpenseByTag")
		return []model.Expense{}, err
	}
	return expenses, nil
}
