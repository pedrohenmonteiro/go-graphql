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

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()

	stmt, err := c.db.Prepare("insert into categories (id, name, description) values (?, ?, ?)")
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

func (c *Category) FindAll() ([]Category, error) {

	rows, err := c.db.Query("select * from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
