package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	config, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong please try again later.")
	}
	config.MaxConns = 10 // the amx of opened channels when the users hit the request
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong please try again later.")
	}
	createTables()

}

func createTables() {
	/// USERS TABLE
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS  users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(300) NOT NULL UNIQUE,
		password VARCHAR(300) NOT NULL
		
	);
	
	`
	_, err := DB.Exec(context.Background(), createUsersTable)
	if err != nil {
		fmt.Println("Something went wrong creating table users.")
	}
	/////////REFRESH TABLE
	createRefreshTokensTable := `
	CREATE TABLE IF NOT EXISTS  refresh_tokens (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users ON DELETE CASCADE,
		token TEXT NOT NULL,
		expire_date TIMESTAMPTZ NOT NULL,
		created_at TIMESTAMPTZ NOT NULL
	);
	
	`
	_, err = DB.Exec(context.Background(), createRefreshTokensTable)
	fmt.Println(err, "EERR")
	if err != nil {
		fmt.Println("Something went wrong creating table users.")
	}

	//// EVENTS TABLE
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TIMESTAMPTZ,
		user_id INT REFERENCES users ON DELETE CASCADE

	);
	`
	_, err = DB.Exec(context.Background(), createEventsTable)
	fmt.Println("--------------", err)
	if err != nil {
		fmt.Println("Something went wrong.")
	}
}
