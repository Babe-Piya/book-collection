package service

import (
	"context"
	"errors"
	"testing"

	mock_repositories "github/Babe-piya/book-collection/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteBookCollectionByIDWhenDeleteSuccessShouldReturnSuccess(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	repo.EXPECT().DeleteBookCollectionByID(context.TODO(), mockID).
		Return(nil)

	s := NewBookCollectionService(repo)

	expected := DeleteBookCollectionResponse{
		Code:    0,
		Message: "Success",
	}
	actual, err := s.DeleteBookCollectionByID(context.TODO(), mockID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestDeleteBookCollectionByIDWhenDeleteFailShouldReturnError(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	mockError := errors.New("mock error")
	repo.EXPECT().DeleteBookCollectionByID(context.TODO(), mockID).
		Return(mockError)

	s := NewBookCollectionService(repo)

	actual, err := s.DeleteBookCollectionByID(context.TODO(), mockID)

	assert.ErrorIs(t, err, mockError)
	assert.Empty(t, actual)
}
