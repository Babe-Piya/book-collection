package repositories

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
)

const (
	mockID                                 = 1
	mockBookName                           = "mock-book-name"
	mockType                               = "mock-type"
	mockVolume                             = 1
	mockPrice                              = 20.00
	sqlGetBookCollectionExpected           = `SELECT * FROM "book_collection" ORDER BY id DESC`
	sqlGetBookCollectionWithFilterExpected = `SELECT * FROM "book_collection" WHERE id = $1 AND book_name = $2 AND type = $3 AND volume = $4 AND price = $5  ORDER BY id DESC`
)

func TestGetBookCollectionByFilterWhenHaveDataShouldReturnData(t *testing.T) {
	mockDB, sqlMock, err := sqlmock.New()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	now := time.Now()
	mockRows := sqlmock.NewRows([]string{"id", "book_name", "type", "volume", "price", "created_at", "updated_at"}).
		AddRow(mockID, mockBookName, mockType, mockVolume, mockPrice, now, now)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sqlGetBookCollectionExpected)).WillReturnRows(mockRows)

	repo := NewBookCollection(gormDB)

	actual, err := repo.GetBookCollectionByFilter(context.Background(), BookCollection{})

	expected := []BookCollection{{
		ID:        mockID,
		BookName:  mockBookName,
		Type:      mockType,
		Volume:    mockVolume,
		Price:     mockPrice,
		CreatedAt: now,
		UpdatedAt: now,
	}}
	assert.NoError(t, err)
	assert.Equal(t, actual, expected)
}

func TestGetBookCollectionByFilterWhenHaveFilterShouldReturnData(t *testing.T) {
	mockDB, sqlMock, err := sqlmock.New()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	now := time.Now()
	mockRows := sqlmock.NewRows([]string{"id", "book_name", "type", "volume", "price", "created_at", "updated_at"}).
		AddRow(mockID, mockBookName, mockType, mockVolume, mockPrice, now, now)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sqlGetBookCollectionWithFilterExpected)).WillReturnRows(mockRows)

	repo := NewBookCollection(gormDB)

	filter := BookCollection{
		ID:       mockID,
		BookName: mockBookName,
		Type:     mockType,
		Volume:   mockVolume,
		Price:    mockPrice,
	}
	actual, err := repo.GetBookCollectionByFilter(context.Background(), filter)

	expected := []BookCollection{{
		ID:        mockID,
		BookName:  mockBookName,
		Type:      mockType,
		Volume:    mockVolume,
		Price:     mockPrice,
		CreatedAt: now,
		UpdatedAt: now,
	}}
	assert.NoError(t, err)
	assert.Equal(t, actual, expected)
}

func TestGetBookCollectionByFilterWhenErrorShouldReturnError(t *testing.T) {
	mockDB, sqlMock, err := sqlmock.New()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	sqlMock.ExpectQuery(regexp.QuoteMeta(sqlGetBookCollectionExpected)).WillReturnError(sql.ErrNoRows)

	repo := NewBookCollection(gormDB)

	actual, err := repo.GetBookCollectionByFilter(context.Background(), BookCollection{})

	assert.ErrorIs(t, err, sql.ErrNoRows)
	assert.Empty(t, actual)
}
