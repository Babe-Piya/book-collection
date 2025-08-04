package service

import (
	"context"
	"log/slog"
)

type DeleteBookCollectionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *bookCollectionService) DeleteBookCollectionByID(ctx context.Context, id int) (DeleteBookCollectionResponse, error) {
	slog.Info("delete book collection by id")
	err := s.BookCollectionRepo.DeleteBookCollectionByID(ctx, id)
	if err != nil {
		slog.Error("delete book collection by id", slog.String("error", err.Error()))

		return DeleteBookCollectionResponse{}, err
	}

	return DeleteBookCollectionResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
