package main

import (
	"log"
	"net/http"

	"github.com/SajibKC/bookapi/api/gen/github.com/SajibKC/bookapi/gen/book/v1/bookv1connect"
	"github.com/SajibKC/bookapi/internal/db"
	"github.com/SajibKC/bookapi/internal/models"
	"github.com/SajibKC/bookapi/internal/service"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.Book{})

	mux := http.NewServeMux()

	bookService := &service.BookServiceServer{}
	path, handler := bookv1connect.NewBookServiceHandler(bookService)
	mux.Handle(path, handler)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
