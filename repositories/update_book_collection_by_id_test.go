package repositories

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	sqlUpdateBookCollectionByIDExpected = `UPDATE "book_collection" SET "book_name"=$1,"type"=$2,"volume"=$3,"price"=$4,"updated_at"=$5 WHERE "id" = $6`
)

func TestUpdateBookCollectionByFilterWhenUpdateDataShouldSuccess(t *testing.T) {
	mockDB, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta(sqlUpdateBookCollectionByIDExpected)).
		WithArgs(mockBookName, mockType, mockVolume, mockPrice, sqlmock.AnyArg(), mockID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	sqlMock.ExpectCommit()

	repo := NewBookCollection(gormDB)

	errExpected := repo.UpdateBookCollectionByID(context.TODO(), BookCollection{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	})

	assert.NoError(t, errExpected)
}

func TestUpdateBookCollectionByFilterWhenUpdateFailShouldFail(t *testing.T) {
	mockDB, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	mockErr := errors.New("mock error")
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta(sqlUpdateBookCollectionByIDExpected)).
		WithArgs(mockBookName, mockType, mockVolume, mockPrice, sqlmock.AnyArg(), mockID).
		WillReturnError(mockErr)
	sqlMock.ExpectRollback()

	repo := NewBookCollection(gormDB)

	errExpected := repo.UpdateBookCollectionByID(context.TODO(), BookCollection{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	})

	assert.ErrorIs(t, errExpected, mockErr)
}
