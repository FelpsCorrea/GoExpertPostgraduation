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

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO category (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{db: c.db, ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var category Category
	err := c.db.QueryRow("SELECT category.id, category.name, category.description FROM category JOIN course ON category.id = course.category_id WHERE course.id = $1", courseID).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}
