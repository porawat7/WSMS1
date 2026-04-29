package repository

import (
	"backend/domain"
	"database/sql"
)

type CourseRepository interface {
	GetAll() ([]domain.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) GetAll() ([]domain.Course, error) {
	rows, err := r.db.Query("SELECT id, name, category, price FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []domain.Course

	for rows.Next() {
		var c domain.Course
		err := rows.Scan(&c.ID, &c.Name, &c.Category, &c.Price)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	return courses, nil
}