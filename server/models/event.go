package models

import (
	"database/sql"
	"server/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {

	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"
	if db.DB == nil {
		db.InitDB()
		if db.DB == nil {
			return nil, nil
		}
	}
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	// Handle the case where no rows are found
	if err == sql.ErrNoRows {
		return nil, err // Return the error if no event is found
	}

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) UpdateEventById() error {
	query := "UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	if err != nil {
		return err
	}
	return nil
}

func DeleteEventById(id int64) error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

func (e Event) Register(userId int64) error {

	query := "INSERT INTO registration (event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	if err != nil {
		return err
	}
	return nil
}

func (e Event) UnRegister(userId int64) error {

	query := "DELETE FROM registration WHERE event_id=? and user_id=?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	if err != nil {
		return err
	}
	return nil
}
