package service

import (
	"context"

	"github/Babe-piya/book-collection/repositories"
)

type BookCollectionService interface {
	CreateBookCollection(ctx context.Context, req BookCollectionRequest) (BookCollectionResponse, error)
	GetBookCollectionByFilter(ctx context.Context, req GetBookCollection) (GetBookCollectionResponse, error)
	UpdateBookCollectionByID(ctx context.Context, req UpdateBookCollectionRequest) (UpdateBookCollectionResponse, error)
	DeleteBookCollectionByID(ctx context.Context, id int) (DeleteBookCollectionResponse, error)
}

type bookCollectionService struct {
	BookCollectionRepo repositories.BookCollectionRepo
}

func NewBookCollectionService(bookCollectionRepo repositories.BookCollectionRepo) BookCollectionService {
	return &bookCollectionService{
		BookCollectionRepo: bookCollectionRepo,
	}
}
