package db

import (
	"log"
	"time"

	"github.com/Random7-JF/xpense-tracker/model"
)

func (s *Sqlite) CheckIfUserExists(username string) bool {
	result, err := s.Db.Query(s.Sql["sql/users/checkForUser.sql"], username)
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
	_, err := s.Db.Exec(s.Sql["sql/users/createUser.sql"], username, HashPassword(password), email, time.Now())
	if err != nil {
		log.Println("exec error for createUser: ", err)
		return err
	}

	return nil
}

func (s *Sqlite) AuthUser(username string, password string) bool {
	result := s.Db.QueryRow(s.Sql["sql/users/getUserPass.sql"], username)
	if result.Err() != nil {
		log.Println("Query error in AuthUser", result.Err())
		return false
	}

	var hash string
	result.Scan(&hash)
	return ComparePassword(hash, password)
}

func (s *Sqlite) GetUserId(username string) string {
	var userId string
	result := s.Db.QueryRow(s.Sql["sql/users/getUserId.sql"], username)
	if result.Err() != nil {
		log.Println("Queryrow error in getuserid", result.Err())
		return ""
	}
	err := result.Scan(&userId)
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
