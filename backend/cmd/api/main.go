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
	// 🔥 connect DB
	connStr := "host=localhost port=5432 user=course_user password=123 dbname=course_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// init layer
	repo := repository.NewCourseRepository(db)
	usecase := usecase.NewCourseUsecase(repo)
	handler := handler.NewCourseHandler(usecase)

	// route
	http.HandleFunc("/api/courses", handler.GetCourses)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}