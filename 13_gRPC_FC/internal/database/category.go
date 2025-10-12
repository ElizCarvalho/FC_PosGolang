package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	stmt, err := c.db.Prepare("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)")
	if err != nil {
		return Category{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) List() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) GetByID(id string) (Category, error) {
	var category Category
	err := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}

// CreateMultiple cria múltiplas categorias simuladas para demonstração do streaming
func (c *Category) CreateMultiple(baseName, baseDescription string, count int) ([]Category, error) {
	var categories []Category

	// Simula criação de múltiplas categorias
	for i := 0; i < count; i++ {
		category, err := c.Create(
			fmt.Sprintf("%s %d", baseName, i+1),
			fmt.Sprintf("%s - Categoria %d", baseDescription, i+1),
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
