package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"db-go.com/api/db"
	"github.com/jackc/pgx/v5"
)

// for binding JSON making sure the incoming data matches with the struct Event
type Event struct {
	ID          int64     `json:"id"          db:"id"`
	Name        string    `json:"name"        db:"name"          binding:"required"`
	Description string    `json:"description" db:"description"   binding:"required"`
	Location    string    `json:"location"    db:"location"      binding:"required"`
	DateTime    time.Time `json:"dateTime"    db:"datetime"`
	UserID      int64     `json:"userId"      db:"user_id"`
}

// var events []Event = []Event{} // the same as bellow but with more redundancy
var events = []Event{} // initialize as empty slice

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	err := db.DB.QueryRow(context.Background(), query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)
	fmt.Println(err, "-x-x-x-x--x-x--x-x-x-x-x-")
	if err != nil {
		return errors.New("Somethign went wrong saving the data.")
	}

	events = append(events, *e)
	return nil
}

func GetEvent(id int64) (*Event, error) {
	query := `
		 SELECT id, name, description, location, datetime, user_id FROM events WHERE id=$1
	`
	row, _ := db.DB.Query(context.Background(), query, id)
	event, err := pgx.CollectOneRow[Event](row, pgx.RowToStructByName[Event])
	fmt.Println(err)
	if err != nil {
		return &event, errors.New("Something went wront on getAll")
	}
	return &event, nil
}
func (body *Event) UpdateEvent() error {
	query := `
		UPDATE events 
		SET name=$1, description=$2, location=$3, datetime=$4
		WHERE id=$5
	`
	_, err := db.DB.Exec(context.Background(), query, body.Name, body.Description, body.Location, body.DateTime, body.ID)

	if err != nil {
		return errors.New("Something went wrong updating the event." + err.Error())
	}
	return nil
}
func GetAllEvents(userId int64) ([]Event, error) {
	fmt.Println(userId)
	query := `
		SELECT * FROM events 
		WHERE user_id=$1
	`
	rows, _ := db.DB.Query(context.Background(), query, userId)
	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[Event])
	if err != nil {
		return nil, errors.New("Something went wront on getAll")
	}
	return events, nil
}

func (event *Event) DeleteEvent() error {
	query := `
		DELETE FROM events where id=$1
	`
	_, err := db.DB.Exec(context.Background(), query, event.ID)

	if err != nil {
		return errors.New("Error when deleting.")
	}
	return nil
}
