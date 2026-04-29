package usecase

import (
	"workspaces/WSMS-final/domain"
)

type courseUsecase struct {
	repo domain.CourseRepository
}

// NewCourseUsecase creates a new instance of courseUsecase.
func NewCourseUsecase(repo domain.CourseRepository) domain.CourseUsecase {
	return &courseUsecase{repo: repo}
}

func (u *courseUsecase) FetchAllCourses() ([]domain.Course, error) {
	return u.repo.GetAllCourses()
}

func (u *courseUsecase) FetchCourseDetails(id int) (domain.Course, error) {
	return u.repo.GetCourseByID(id)
}