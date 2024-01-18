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

func (c *Course) Create(name string, description string, categoryID string) (*Course, error) {
	id := uuid.New().String()

	stmt, err := c.db.Prepare("insert into courses (id, name, description, category_id) values (?, ?, ?, ?)")
	if err != nil {
		return &Course{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description, categoryID)

	if err != nil {
		return &Course{}, nil
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil

}

func (c *Course) FindAll() ([]Course, error) {

	rows, err := c.db.Query("select * from courses")
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
