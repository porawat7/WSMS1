package repository

import (
	"database/sql"
	"workspaces/WSMS-final/domain"
)

type sqliteCourseRepository struct {
	db *sql.DB
}

// NewSQLiteCourseRepository creates a new instance of sqliteCourseRepository.
func NewSQLiteCourseRepository(db *sql.DB) domain.CourseRepository {
	return &sqliteCourseRepository{db: db}
}

func (r *sqliteCourseRepository) GetAllCourses() ([]domain.Course, error) {
	rows, err := r.db.Query("SELECT id, name, category, price, description, platform, link, startDate, time FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		var course domain.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Category, &course.Price, &course.Description, &course.Platform, &course.Link, &course.StartDate, &course.Time); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (r *sqliteCourseRepository) GetCourseByID(id int) (domain.Course, error) {
	var course domain.Course
	err := r.db.QueryRow("SELECT id, name, category, price, description, platform, link, startDate, time FROM courses WHERE id = ?", id).
		Scan(&course.ID, &course.Name, &course.Category, &course.Price, &course.Description, &course.Platform, &course.Link, &course.StartDate, &course.Time)
	if err != nil {
		return domain.Course{}, err
	}

	return course, nil
}