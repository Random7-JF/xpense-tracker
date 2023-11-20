package db

import (
	"log"
	"time"

	"github.com/Random7-JF/xpense-tracker/model"
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

	_, err = s.Db.Exec(query, username, HashPassword(password), email, time.Now())
	if err != nil {
		log.Println("exec error for createUser: ", err)
		return err
	}

	return nil
}

func (s *Sqlite) AuthUser(username string, password string) bool {
	query, err := ReadSQL("users/getUserPass.sql")
	if err != nil {
		log.Panicln("Readsql error for AuthUser", err)
		return false
	}

	result := s.Db.QueryRow(query, username)
	if result.Err() != nil {
		log.Println("Query error in AuthUser", err)
		return false
	}

	var hash string
	result.Scan(&hash)
	return ComparePassword(hash, password)
}

func (s *Sqlite) GetUserId(username string) string {
	var userId string
	query, err := ReadSQL("users/getUserId.sql")
	if err != nil {
		log.Println("Query error in Getuserid", err)
		return ""
	}
	result := s.Db.QueryRow(query, username)
	if result.Err() != nil {
		log.Println("Queryrow error in getuserid", err)
		return ""
	}
	err = result.Scan(&userId)
	if err != nil {
		log.Println("scan error", err)
		return ""
	}
	return userId
}

func (s *Sqlite) GetUsers() []model.User {
	var users []model.User
	result, err := s.Db.Query(s.Sql["sql/users/getUsers.sql"])
	if err != nil {
		log.Printf("Error in query of users: %s", err)
	}
	for result.Next() {
		var user model.User
		err := result.Scan(&user.Id, &user.Username, &user.Hasedpassword, &user.Email, &user.Creationdate, &user.Lastlogin)
		if err != nil {
			log.Printf("Scan Error: %s", err)
		}
		users = append(users, user)
	}
	return users
}
