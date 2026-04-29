package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	handler "backend/delivery/http"
	"backend/repository"
	"backend/usecase"
)

func main() {
	// connect DB
	connStr := "host=localhost port=5432 user=course_user password=123 dbname=course_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// init layers
	repo := repository.NewCourseRepository(db)
	usecase := usecase.NewCourseUsecase(repo)
	handler := handler.NewCourseHandler(usecase)

	// ===== ROUTES =====

	// GET all courses
	http.HandleFunc("/api/courses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler.GetCourses(w, r)
			return
		}

		if r.Method == "POST" {
			handler.CreateCourse(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	// GET by id, PUT, DELETE
	http.HandleFunc("/api/courses/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			handler.GetCourseByID(w, r)
			return
		}

		if r.Method == "PUT" {
			handler.UpdateCourse(w, r)
			return
		}

		if r.Method == "DELETE" {
			handler.DeleteCourse(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}