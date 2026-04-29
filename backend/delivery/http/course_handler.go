package http

import (
	"encoding/json"
	"net/http"

	"backend/usecase"
)

type CourseHandler struct {
	usecase usecase.CourseUsecase
}

func NewCourseHandler(u usecase.CourseUsecase) *CourseHandler {
	return &CourseHandler{usecase: u}
}

func (h *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, _ := h.usecase.GetCourses()
	json.NewEncoder(w).Encode(courses)
}