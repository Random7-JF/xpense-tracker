package db

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/Random7-JF/xpense-tracker/model"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetExpense(userId string) ([]model.Expense, error)
	GetExpenseByID(expenseID string) (model.Expense, error)
	GetExpenseByFreq(freq string, userid string) ([]model.Expense, error)
	GetExpenseBySearch(search string, userid string) ([]model.Expense, error)
	GetExpenseByTag(tag string, userid string) ([]model.Expense, error)
	AddExpense(e model.Expense) error
	RemoveExpense(expenseId int) error
	UpdateExpenseById(expense model.Expense) error
	ExpenseFill(userid string) error
	GetExpenseCountByUser(username string) (int, error)

	CheckIfUserExists(username string) bool
	CreateUser(username string, password string, email string) error
	AuthUser(username string, password string) bool
	GetUserId(username string) string
	GetUsers() []model.User
}

type Sqlite struct {
	Db    *sql.DB
	Sql   map[string]string
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
	s.Sql, _ = readSQL("sql/")
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

func readSQL(folderPath string) (map[string]string, error) {
	sqlfiles := make(map[string]string)
	err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %s", path, err)
		}
		if d.IsDir() {
			return nil
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error Reading file %s: %s", path, err)
		}
		sqlfiles[path] = string(contents)
		return nil
	})

	if err != nil {
		log.Printf("Error walking %s: %s", folderPath, err)
	}

	return sqlfiles, nil
}
