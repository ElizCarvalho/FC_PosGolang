package database

import (
	"database/sql"

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

// Create cria uma nova categoria
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

// List retorna todas as categorias
func (c *Category) List() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories ORDER BY name")
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

// GetByID busca uma categoria por ID
func (c *Category) GetByID(id string) (Category, error) {
	var category Category
	err := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}

// Update atualiza uma categoria existente
func (c *Category) Update(id, name, description string) error {
	stmt, err := c.db.Prepare("UPDATE categories SET name = ?, description = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, id)
	return err
}

// Delete remove uma categoria
func (c *Category) Delete(id string) error {
	stmt, err := c.db.Prepare("DELETE FROM categories WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
