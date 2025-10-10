package database

import "database/sql"

type Course struct {
	db          *sql.DB
	ID          int
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}
