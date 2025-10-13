package service

import (
	"context"
	"fmt"

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

	if err := db.DB.Create(&book).Error; err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	res := connect.NewResponse(&bookv1.CreateBookResponse{
		Book: &bookv1.Book{
			Id:     fmt.Sprintf("%d", book.ID),
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
	if err := db.DB.Find(&books).Error; err != nil {
		return nil, fmt.Errorf("failed to list books: %w", err)
	}

	respBooks := make([]*bookv1.Book, len(books))
	for i, b := range books {
		respBooks[i] = &bookv1.Book{
			Id:     fmt.Sprintf("%d", b.ID),
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

func (s *BookServiceServer) GetBook(
	ctx context.Context,
	req *connect.Request[bookv1.GetBookRequest],
) (*connect.Response[bookv1.GetBookResponse], error) {

	var book models.Book
	if err := db.DB.First(&book, req.Msg.Id).Error; err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}

	res := connect.NewResponse(&bookv1.GetBookResponse{
		Book: &bookv1.Book{
			Id:     fmt.Sprintf("%d", book.ID),
			Title:  book.Title,
			Author: book.Author,
			Price:  book.Price,
		},
	})
	return res, nil
}

func (s *BookServiceServer) UpdateBook(
	ctx context.Context,
	req *connect.Request[bookv1.UpdateBookRequest],
) (*connect.Response[bookv1.UpdateBookResponse], error) {

	var book models.Book
	if err := db.DB.First(&book, req.Msg.Id).Error; err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}

	book.Title = req.Msg.Title
	book.Author = req.Msg.Author
	book.Price = req.Msg.Price

	if err := db.DB.Save(&book).Error; err != nil {
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	res := connect.NewResponse(&bookv1.UpdateBookResponse{
		Book: &bookv1.Book{
			Id:     fmt.Sprintf("%d", book.ID),
			Title:  book.Title,
			Author: book.Author,
			Price:  book.Price,
		},
	})
	return res, nil
}

func (s *BookServiceServer) DeleteBook(
	ctx context.Context,
	req *connect.Request[bookv1.DeleteBookRequest],
) (*connect.Response[bookv1.DeleteBookResponse], error) {

	if err := db.DB.Delete(&models.Book{}, req.Msg.Id).Error; err != nil {
		return nil, fmt.Errorf("failed to delete book: %w", err)
	}

	res := connect.NewResponse(&bookv1.DeleteBookResponse{
		Message: "Book deleted successfully",
	})
	return res, nil
}
