package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type BookCollectionRepo interface {
	CreateBookCollection(ctx context.Context, bookCollection *BookCollection) (*BookCollection, error)
}

type bookCollectionRepo struct {
	DB *gorm.DB
}

func NewBookCollection(db *gorm.DB) BookCollectionRepo {
	return &bookCollectionRepo{
		DB: db,
	}
}

type BookCollection struct {
	ID        int
	BookName  string
	Type      string
	Volume    int
	Price     float64
	CreatedAt time.Time
}

func (r *BookCollection) TableName() string {
	return "book_collection"
}
