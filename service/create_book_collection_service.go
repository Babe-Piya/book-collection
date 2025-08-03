package service

import (
	"context"
	"log/slog"

	"github/Babe-piya/book-collection/repositories"
)

type BookCollectionRequest struct {
	BookName string  `json:"book_name"`
	Type     string  `json:"type"`
	Volume   int     `json:"volume"`
	Price    float64 `json:"price"`
}

type BookCollectionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ID      int
}

func (s *bookCollectionService) CreateBookCollection(ctx context.Context, req BookCollectionRequest) (
	BookCollectionResponse, error) {
	slog.Info("create book collection")
	bookCollect, err := s.BookCollectionRepo.CreateBookCollection(ctx, &repositories.BookCollection{
		BookName: req.BookName,
		Type:     req.Type,
		Volume:   req.Volume,
		Price:    req.Price,
	})
	if err != nil {
		slog.Error("error create book collection", slog.String("error", err.Error()))

		return BookCollectionResponse{}, err
	}

	return BookCollectionResponse{
		Code:    0,
		Message: "Success",
		ID:      bookCollect.ID,
	}, nil
}
