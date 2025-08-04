package repositories

import "context"

func (repo *bookCollectionRepo) UpdateBookCollectionByID(ctx context.Context, bookCollections BookCollection) error {
	return repo.DB.WithContext(ctx).Updates(&bookCollections).Error
}
