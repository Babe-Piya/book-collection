package repositories

import "context"

func (repo *bookCollectionRepo) CreateBookCollection(ctx context.Context, bookCollections *BookCollection) (
	*BookCollection, error) {
	err := repo.DB.WithContext(ctx).Create(&bookCollections).Error

	return bookCollections, err
}
