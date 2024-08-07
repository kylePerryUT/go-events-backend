package models

import (
	"time"

	"example.com/example/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (event Event) Save() error {
	// Save the event to the database
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	event.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
	select * from events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			panic(err)
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int64) (Event, error) {
	query := `
	select * from events where id = ?
	`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		panic(err)
	}
	return event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `
		DELETE FROM events
		WHERE id = ?
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}