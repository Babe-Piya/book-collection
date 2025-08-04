package service

import (
	"context"
	"log/slog"

	"github/Babe-piya/book-collection/repositories"
)

type UpdateBookCollectionRequest struct {
	ID       int     `json:"id"`
	BookName string  `json:"book_name"`
	Type     string  `json:"type"`
	Volume   int     `json:"volume"`
	Price    float64 `json:"price"`
}

type UpdateBookCollectionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *bookCollectionService) UpdateBookCollectionByID(ctx context.Context, req UpdateBookCollectionRequest) (
	UpdateBookCollectionResponse, error) {
	slog.Info("update book collection by id")
	err := s.BookCollectionRepo.UpdateBookCollectionByID(ctx, repositories.BookCollection{
		ID:       req.ID,
		BookName: req.BookName,
		Type:     req.Type,
		Volume:   req.Volume,
		Price:    req.Price,
	})
	if err != nil {
		slog.Error("error update book collection by id", slog.String("error", err.Error()))

		return UpdateBookCollectionResponse{}, err
	}

	return UpdateBookCollectionResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
