package service

import (
	"context"

	"connectrpc.com/connect"
	bookv1 "github.com/SajibKC/bookapi/api/gen/github.com/SajibKC/bookapi/gen/book/v1"
	"github.com/SajibKC/bookapi/internal/db"
	"github.com/SajibKC/bookapi/internal/models"
)

type BookServiceServer struct{}

func (s *BookServiceServer) CreateBook(
	ctx context.Context,
	req *connect.Request[bookv1.CreateBookRequest],
) (*connect.Response[bookv1.CreateBookResponse], error) {

	book := models.Book{
		Title:  req.Msg.Title,
		Author: req.Msg.Author,
		Price:  req.Msg.Price,
	}

	db.DB.Create(&book)

	res := connect.NewResponse(&bookv1.CreateBookResponse{
		Book: &bookv1.Book{
			Id:     int64(book.ID),
			Title:  book.Title,
			Author: book.Author,
			Price:  book.Price,
		},
	})
	return res, nil
}

func (s *BookServiceServer) ListBooks(
	ctx context.Context,
	req *connect.Request[bookv1.ListBooksRequest],
) (*connect.Response[bookv1.ListBooksResponse], error) {

	var books []models.Book
	db.DB.Find(&books)

	respBooks := make([]*bookv1.Book, len(books))
	for i, b := range books {
		respBooks[i] = &bookv1.Book{
			Id:     int64(b.ID),
			Title:  b.Title,
			Author: b.Author,
			Price:  b.Price,
		}
	}

	res := connect.NewResponse(&bookv1.ListBooksResponse{
		Books: respBooks,
	})
	return res, nil
}
