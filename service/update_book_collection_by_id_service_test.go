package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github/Babe-piya/book-collection/repositories"

	mock_repositories "github/Babe-piya/book-collection/repositories/mocks"

	"go.uber.org/mock/gomock"
)

func TestUpdateBookCollectionByIDWhenUpdateSuccessShouldReturnSuccess(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	mockBookCollection := repositories.BookCollection{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}
	repo.EXPECT().UpdateBookCollectionByID(context.TODO(), mockBookCollection).
		Return(nil)

	s := NewBookCollectionService(repo)

	expected := UpdateBookCollectionResponse{
		Code:    0,
		Message: "Success",
	}
	actual, err := s.UpdateBookCollectionByID(context.TODO(), UpdateBookCollectionRequest{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	})

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUpdateBookCollectionByIDWhenErrorShouldReturnError(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	mockError := errors.New("mock error")
	mockBookCollection := repositories.BookCollection{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}
	repo.EXPECT().UpdateBookCollectionByID(context.TODO(), mockBookCollection).
		Return(mockError)

	s := NewBookCollectionService(repo)

	actual, err := s.UpdateBookCollectionByID(context.TODO(), UpdateBookCollectionRequest{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	})

	assert.ErrorIs(t, mockError, err)
	assert.Empty(t, actual)
}
