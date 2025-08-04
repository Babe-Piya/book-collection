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
	sqlInsertBookCollectionExpected = `INSERT INTO "book_collection" ("book_name","type","volume","price","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`
)

func TestCreateBookCollectionWhenSuccessShouldReturnData(t *testing.T) {
	mockDB, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(regexp.QuoteMeta(sqlInsertBookCollectionExpected)).
		WithArgs(mockBookName, mockType, mockVolume, mockPrice, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockID))
	sqlMock.ExpectCommit()

	repo := NewBookCollection(gormDB)

	expected := &BookCollection{
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}
	actual, err := repo.CreateBookCollection(context.TODO(), expected)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCreateBookCollectionWhenFailShouldReturnError(t *testing.T) {
	mockDB, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	mockError := errors.New("mock error")
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(regexp.QuoteMeta(sqlInsertBookCollectionExpected)).
		WithArgs(mockBookName, mockType, mockVolume, mockPrice, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(mockError)
	sqlMock.ExpectRollback()

	repo := NewBookCollection(gormDB)

	expected := &BookCollection{
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}
	actual, errExpected := repo.CreateBookCollection(context.TODO(), expected)

	assert.ErrorIs(t, mockError, errExpected)
	assert.Equal(t, expected, actual)
}
