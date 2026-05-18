package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"db-go.com/api/db"
)

type Session struct {
	ID           int64     `json:"id" db:"id"`
	Token        string    `json:"token" db:"token"`
	User_id      int64     `json:"user_id" db:"user_id"`
	Expired_date time.Time `json:"expired_at" db:"expired_date"`
	Created_at   time.Time `json:"created_at" db:"created_at"`
}

func Save(userId int64, expiredDate time.Time, token string) error {
	created_at := time.Now()
	query := `
		INSERT INTO refresh_tokens (token,user_id, expire_date,created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := db.DB.Exec(context.Background(), query, token, userId, expiredDate, created_at)
	fmt.Println(err)
	if err != nil {
		return errors.New("Sommething went wrong inserting Session")
	}
	return nil
}

func Delete(token string) error {
	query := `
		DELETE FROM refresh_tokens
		WHERE token=$1
	`
	_, err := db.DB.Exec(context.Background(), query, token)
	fmt.Println(err)
	if err != nil {
		return errors.New("Sommething went wrong inserting Session")
	}
	return nil
}
