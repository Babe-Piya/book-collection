package repositories

import "context"

func (repo *bookCollectionRepo) GetBookCollectionByFilter(ctx context.Context, filter BookCollection) (
	[]BookCollection, error) {
	query := repo.DB.WithContext(ctx).Order("id DESC")

	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.BookName != "" {
		query = query.Where("book_name = ?", filter.BookName)
	}
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}
	if filter.Volume != 0 {
		query = query.Where("volume = ?", filter.Volume)
	}
	if filter.Price != 0 {
		query = query.Where("price = ?", filter.Price)
	}

	var result []BookCollection
	err := query.Find(&result).Error

	return result, err
}
