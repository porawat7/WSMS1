package usecase

import (
	"backend/domain"
	"backend/repository"
)

type CourseUsecase interface {
	GetCourses() ([]domain.Course, error)
	GetCourseByID(id int) (*domain.Course, error)
	CreateCourse(c domain.Course) error
	UpdateCourse(id int, c domain.Course) error
	DeleteCourse(id int) error
}

type courseUsecase struct {
	repo repository.CourseRepository
}

func NewCourseUsecase(r repository.CourseRepository) CourseUsecase {
	return &courseUsecase{repo: r}
}

func (u *courseUsecase) GetCourses() ([]domain.Course, error) {
	return u.repo.GetAll()
}

func (u *courseUsecase) GetCourseByID(id int) (*domain.Course, error) {
	return u.repo.GetByID(id)
}

func (u *courseUsecase) CreateCourse(c domain.Course) error {
	return u.repo.Create(c)
}

func (u *courseUsecase) UpdateCourse(id int, c domain.Course) error {
	return u.repo.Update(id, c)
}

func (u *courseUsecase) DeleteCourse(id int) error {
	return u.repo.Delete(id)
}