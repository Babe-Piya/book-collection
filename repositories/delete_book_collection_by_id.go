package repositories

import "context"

func (repo *bookCollectionRepo) DeleteBookCollectionByID(ctx context.Context, id int) error {
	return repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&BookCollection{}).Error
}
