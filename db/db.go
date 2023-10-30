package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Random7-JF/xpense-tracker/model"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetExpense(userId string) ([]model.Expense, error)
	AddExpense(e model.Expense) error
	RemoveExpense(expenseId int) error

	CheckIfUserExists(username string) bool
	CreateUser(username string, password string, email string) error
	AuthUser(username string, password string) bool
	GetUserId(username string) string
}

type Sqlite struct {
	Db    *sql.DB
	Error error
}

func ConnectSqliteDb(host string) *Sqlite {
	var sqlDB Sqlite
	sqlDB.Db, sqlDB.Error = sql.Open("sqlite3", host)
	if sqlDB.Error != nil {
		log.Fatalln("No Database connection. ", sqlDB.Error)
	}
	sqlDB.InitDb()
	return &sqlDB
}

func ReadSQL(file string) (string, error) {
	filename := fmt.Sprintf("sql/%s", file)
	sql, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error in reading SQL: ", err)
		return "", err
	}
	return string(sql), nil
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash)
}

func ComparePassword(hashpassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func (s *Sqlite) InitDb() error {
	query, err := ReadSQL("createTables.sql")
	if err != nil {
		log.Println("Error in InitDb: ", err)
		return err
	}
	result, err := s.Db.Exec(query)
	if err != nil {
		log.Println("Error in InitDb: ", err)
		return err
	}
	log.Println(result.RowsAffected())
	return nil
}
