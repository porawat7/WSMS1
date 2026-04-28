package main

import (
	"database/sql"
	"log"
	"net/http"

	"workspaces/oop102/backend/repository"
	"workspaces/oop102/coffee-shop/usecase"
	"workspaces/oop102/backend/delivery/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// เปิด SQLite connection
	db, err := sql.Open("sqlite3", "coffee_shop.db")
	if err != nil {
		log.Fatal("เปิด SQLite ไม่ได้:", err)
	}
	defer db.Close()

	// สร้าง repositories
	coffeeRepo := repository.NewSQLiteCoffeeRepo(db)
	orderRepo := repository.NewSQLiteOrderRepo(db)

	// inject repositories เข้า usecase
	orderUC := usecase.NewOrderUseCase(coffeeRepo, orderRepo)
	// สมมติว่ามี usecase สำหรับ summary
	// summaryUC := usecase.NewSummaryUseCase(orderRepo)

	// inject usecase เข้า HTTP handler
	coffeeHandler := http.NewCoffeeHandler(orderUC)
	orderHandler := http.NewOrderHandler(orderUC)
	// summaryHandler := http.NewSummaryHandler(summaryUC)

	// wiring routes
	http.Handle("/coffees", coffeeHandler)
	http.Handle("/orders", orderHandler)
	// http.Handle("/summary", summaryHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
