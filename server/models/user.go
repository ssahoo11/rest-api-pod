package models

import (
	"errors"
	"server/db"
	"server/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {

	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hash, err := utils.GetHashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hash)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)
	var password string
	err := row.Scan(&u.ID, &password)
	if err != nil {
		return errors.New("credentials invalid")
	}

	matched := utils.ComparePassword(u.Password, password)
	if !matched {
		return errors.New("credentials invalid")
	}
	return nil
}

// func (u User) Exists() bool {
// 	query := "SELECT * FROM users WHERE email=? and id=?"
// 	row := db.DB.QueryRow(query, u.Email, u.ID)
// 	var password string
// 	err := row.Scan(&u.ID, &password)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }
