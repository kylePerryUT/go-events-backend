package models

import "example.com/example/db"

type USER struct {
	ID       int64
	NAME     string `binding:"required"`
	EMAIL    string `binding:"required"`
	PASSWORD string `binding:"required"`
}

func (user USER) Save() error {
	// Save the user to the database
	query := `
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?)
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.NAME, user.EMAIL, user.PASSWORD)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.ID = id
	return nil
}