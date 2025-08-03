package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github/Babe-piya/book-collection/repositories"
	mock_repositories "github/Babe-piya/book-collection/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	mockID       = 1
	mockBookName = "mock-book-name"
	mockType     = "mock-type"
	mockVolume   = 1
	mockPrice    = 20
)

func TestGetBookCollectionByFilterWhenHaveDataShouldReturnData(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	now := time.Now()
	mockBookCollection := []repositories.BookCollection{{
		ID:        mockID,
		BookName:  mockBookName,
		Type:      mockType,
		Volume:    mockVolume,
		Price:     mockPrice,
		CreatedAt: now,
	}}

	repo.EXPECT().GetBookCollectionByFilter(context.TODO(), repositories.BookCollection{ID: mockID}).
		Return(mockBookCollection, nil)

	s := NewBookCollectionService(repo)

	expected := GetBookCollectionResponse{
		Code:    0,
		Message: "Success",
		Data: []GetBookCollection{{
			ID:       mockID,
			BookName: mockBookName,
			Type:     mockType,
			Volume:   mockVolume,
			Price:    mockPrice,
		}},
	}
	actual, err := s.GetBookCollectionByFilter(context.TODO(), GetBookCollection{ID: mockID})

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetBookCollectionByFilterWhenErrorShouldReturnFail(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	mockError := errors.New("mock error")
	repo.EXPECT().GetBookCollectionByFilter(context.TODO(), repositories.BookCollection{ID: mockID}).
		Return(nil, mockError)

	s := NewBookCollectionService(repo)

	actual, err := s.GetBookCollectionByFilter(context.TODO(), GetBookCollection{ID: mockID})

	assert.ErrorIs(t, err, mockError)
	assert.Empty(t, actual)
}
