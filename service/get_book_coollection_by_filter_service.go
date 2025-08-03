package service

import (
	"context"
	"log/slog"

	"github/Babe-piya/book-collection/repositories"
)

type GetBookCollectionResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []GetBookCollection `json:"data"`
}

type GetBookCollection struct {
	ID       int     `json:"id"`
	BookName string  `json:"book_name"`
	Type     string  `json:"type"`
	Volume   int     `json:"volume"`
	Price    float64 `json:"price"`
}

func (s *bookCollectionService) GetBookCollectionByFilter(ctx context.Context, req GetBookCollection) (
	GetBookCollectionResponse, error) {
	slog.Info("get book collection by filter")
	bookCollects, err := s.BookCollectionRepo.GetBookCollectionByFilter(ctx, repositories.BookCollection{
		ID:       req.ID,
		BookName: req.BookName,
		Type:     req.Type,
		Volume:   req.Volume,
		Price:    req.Price,
	})
	if err != nil {
		slog.Error("error get book collection by filter", err)

		return GetBookCollectionResponse{}, err
	}

	var data []GetBookCollection
	for _, bookCollect := range bookCollects {
		data = append(data, GetBookCollection{
			ID:       bookCollect.ID,
			BookName: bookCollect.BookName,
			Type:     bookCollect.Type,
			Volume:   bookCollect.Volume,
			Price:    bookCollect.Price,
		})
	}

	return GetBookCollectionResponse{
		Code:    0,
		Message: "Success",
		Data:    data,
	}, nil
}
