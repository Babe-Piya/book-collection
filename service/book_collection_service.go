package service

import (
	"context"
	"github/Babe-piya/book-collection/repositories"
)

type BookCollectionService interface {
	CreateBookCollection(ctx context.Context, req BookCollectionRequest) (BookCollectionResponse, error)
}

type bookCollectionService struct {
	BookCollectionRepo repositories.BookCollectionRepo
}

func NewBookCollectionService(bookCollectionRepo repositories.BookCollectionRepo) BookCollectionService {
	return &bookCollectionService{
		BookCollectionRepo: bookCollectionRepo,
	}
}
