package db

import (
	"log"
	"time"
)

func (s *Sqlite) CheckIfUserExists(username string) bool {
	query, err := ReadSQL("users/checkForUser.sql")
	if err != nil {
		log.Println("Readsql error for CheckforUser: ", err)
		return true
	}

	result, err := s.Db.Query(query, username)
	if err != nil {
		log.Println("Checkforuser query error:", err)
		return true
	}

	if result.Next() {
		return true
	}
	return false
}

func (s *Sqlite) CreateUser(username string, password string, email string) error {
	query, err := ReadSQL("users/createUser.sql")
	if err != nil {
		log.Println("Readsql error for createUser: ", err)
		return err
	}

	_, err = s.Db.Exec(query, username, hashPassword(password), email, time.Now())
	if err != nil {
		log.Println("exec error for createUser: ", err)
		return err
	}

	return nil
}
