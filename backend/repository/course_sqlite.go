package repository

import (
	"backend/domain"
	"database/sql"
)

type CourseRepository interface {
	GetAll() ([]domain.Course, error)
	GetByID(id int) (*domain.Course, error)
	Create(c domain.Course) error
	Update(id int, c domain.Course) error
	Delete(id int) error
}

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}

// ================= GET ALL =================
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

// ================= GET BY ID =================
func (r *courseRepository) GetByID(id int) (*domain.Course, error) {
	row := r.db.QueryRow("SELECT id, name, category, price FROM courses WHERE id=$1", id)

	var c domain.Course
	err := row.Scan(&c.ID, &c.Name, &c.Category, &c.Price)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// ================= CREATE =================
func (r *courseRepository) Create(c domain.Course) error {
	return r.db.QueryRow(
		"INSERT INTO courses (name, category, price) VALUES ($1,$2,$3) RETURNING id",
		c.Name, c.Category, c.Price,
	).Scan(&c.ID)
}

// ================= UPDATE =================
func (r *courseRepository) Update(id int, c domain.Course) error {
	_, err := r.db.Exec(
		"UPDATE courses SET name=$1, category=$2, price=$3 WHERE id=$4",
		c.Name, c.Category, c.Price, id,
	)
	return err
}

// ================= DELETE =================
func (r *courseRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM courses WHERE id=$1", id)
	return err
}