package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetExpense()
	AddExpense()
	RemoveExpense()

	CheckIfUserExists(username string) bool
	CreateUser(username string, password string, email string) error
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

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash)
}
