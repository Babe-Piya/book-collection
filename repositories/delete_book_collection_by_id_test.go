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
	sqlDeleteBookCollectionWithIDExpected = `DELETE FROM "book_collection" WHERE id = $1`
)

func TestDeleteBookCollectionByIDWhenDeleteSuccessShouldSuccess(t *testing.T) {
	mockDB, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta(sqlDeleteBookCollectionWithIDExpected)).
		WithArgs(mockID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	sqlMock.ExpectCommit()

	repo := NewBookCollection(gormDB)

	err := repo.DeleteBookCollectionByID(context.Background(), mockID)

	assert.NoError(t, err)
}

func TestDeleteBookCollectionByIDWhenDeleteFailShouldReturnError(t *testing.T) {
	mockDB, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	mockError := errors.New("mock error")
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta(sqlDeleteBookCollectionWithIDExpected)).
		WithArgs(mockID).
		WillReturnError(mockError)
	sqlMock.ExpectRollback()

	repo := NewBookCollection(gormDB)

	err := repo.DeleteBookCollectionByID(context.Background(), mockID)

	assert.ErrorIs(t, err, mockError)
}
