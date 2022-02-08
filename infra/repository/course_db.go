package repository

import (
	"database/sql"
	"main/entity"
)

type CourseMySQLRepository struct {
	DB *sql.DB
}

func (c CourseMySQLRepository) Insert(course entity.Course) error {
	stmt, err := c.DB.Prepare(`insert into courses (id, name, description, status) values (?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	return err
}
