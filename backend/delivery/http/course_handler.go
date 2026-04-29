package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	
	"backend/domain"
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

// GET /api/courses/:id
func (h *CourseHandler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/courses/"):]
	id, _ := strconv.Atoi(idStr)

	course, err := h.usecase.GetCourseByID(id)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}

	json.NewEncoder(w).Encode(course)
}

// POST
func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var c domain.Course
	json.NewDecoder(r.Body).Decode(&c)

	h.usecase.CreateCourse(c)
	w.Write([]byte("created"))
}

// PUT
func (h *CourseHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/courses/"):]
	id, _ := strconv.Atoi(idStr)

	var c domain.Course
	json.NewDecoder(r.Body).Decode(&c)

	h.usecase.UpdateCourse(id, c)
	w.Write([]byte("updated"))
}

// DELETE
func (h *CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/courses/"):]
	id, _ := strconv.Atoi(idStr)

	h.usecase.DeleteCourse(id)
	w.Write([]byte("deleted"))
}