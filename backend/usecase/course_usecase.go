package usecase

import (
	"backend/domain"
	"backend/repository"
)

type CourseUsecase interface {
	GetCourses() ([]domain.Course, error)
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