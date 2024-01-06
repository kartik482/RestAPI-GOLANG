package models

import (
	"example.com/api/db"
	"example.com/api/utils"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.Hash(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()

	return err
}
func (u *User) ValidateCredentials() error {
	query := `SELECT id,password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !passwordIsValid {
		return err
	}
	return nil

}
