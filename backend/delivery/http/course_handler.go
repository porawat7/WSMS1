package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"workspaces/WSMS-final/domain"
)

type CourseHandler struct {
	Usecase domain.CourseUsecase
}

func NewCourseHandler(usecase domain.CourseUsecase) *CourseHandler {
	return &CourseHandler{Usecase: usecase}
}

func (h *CourseHandler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.Usecase.FetchAllCourses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func (h *CourseHandler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	course, err := h.Usecase.FetchCourseDetails(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}