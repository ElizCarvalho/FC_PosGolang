package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryID string) (Course, error) {
	id := uuid.New().String()
	stmt, err := c.db.Prepare("INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return Course{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}
	return Course{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) List() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		err = rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
