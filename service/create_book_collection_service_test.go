package service

import (
	"context"
	"errors"
	"testing"

	"github/Babe-piya/book-collection/repositories"

	mock_repositories "github/Babe-piya/book-collection/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateBookCollectionWhenCreateSuccessShouldReturnSuccess(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	mockBookCollectionResp := repositories.BookCollection{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}
	repo.EXPECT().CreateBookCollection(context.TODO(), &repositories.BookCollection{
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}).
		Return(&mockBookCollectionResp, nil)

	s := NewBookCollectionService(repo)

	expected := BookCollectionResponse{
		Code:    0,
		Message: "Success",
		ID:      mockID,
	}
	actual, err := s.CreateBookCollection(context.TODO(), BookCollectionRequest{
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	})

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCreateBookCollectionWhenCreateFailShouldReturnError(t *testing.T) {
	mock_ctl := gomock.NewController(t)

	repo := mock_repositories.NewMockBookCollectionRepo(mock_ctl)

	mockBookCollection := repositories.BookCollection{
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}
	mockError := errors.New("mock error")
	repo.EXPECT().CreateBookCollection(context.TODO(), &mockBookCollection).
		Return(nil, mockError)

	s := NewBookCollectionService(repo)

	actual, err := s.CreateBookCollection(context.TODO(), BookCollectionRequest{
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	})

	assert.ErrorIs(t, err, mockError)
	assert.Empty(t, actual)
}
