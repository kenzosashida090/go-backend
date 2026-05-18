package models

import (
	"context"
	"errors"
	"fmt"

	"db-go.com/api/db"
	"db-go.com/api/utils"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email,password)
		VALUES ($1,$2)
		RETURNING id
	`
	hashedPassword, err := utils.HashingPassword(u.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	err = db.DB.QueryRow(context.Background(), query, u.Email, hashedPassword).Scan(&u.ID)

	if err != nil {
		return errors.New("Something went wrong saving user")
	}

	return nil
}

func (u *User) GetUser(email string) error {
	query := `
		SELECT password,id FROM users WHERE email=$1
	`
	err := db.DB.QueryRow(context.Background(), query, email).Scan(&u.Password, &u.ID)

	fmt.Println(u.Password, "x0x00x0x0x0x0")
	if err != nil {
		return errors.New("Somethingg went wrong when fetching user.")
	}

	return nil
}
