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

func (c *Course) GetByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = ?", categoryID)
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

func (c *Course) GetByID(id string) (Course, error) {
	var course Course
	err := c.db.QueryRow("SELECT id, name, description, category_id FROM courses WHERE id = ?", id).Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

// CourseWithCategory representa um curso com dados da categoria
type CourseWithCategory struct {
	CourseID            string
	CourseName          string
	CourseDescription   string
	CategoryID          string
	CategoryName        string
	CategoryDescription string
}

// ListWithCategories busca cursos com dados das categorias usando JOIN
func (c *Course) ListWithCategories() ([]CourseWithCategory, error) {
	query := `
		SELECT 
			c.id as course_id,
			c.name as course_name,
			c.description as course_description,
			c.category_id,
			cat.name as category_name,
			cat.description as category_description
		FROM courses c
		JOIN categories cat ON c.category_id = cat.id
		ORDER BY c.name
	`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []CourseWithCategory
	for rows.Next() {
		var course CourseWithCategory
		err := rows.Scan(
			&course.CourseID,
			&course.CourseName,
			&course.CourseDescription,
			&course.CategoryID,
			&course.CategoryName,
			&course.CategoryDescription,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
