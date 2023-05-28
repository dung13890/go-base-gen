package repository

import (
	"context"

	"project-demo/internal/domain"
	"project-demo/pkg/errors"

	"gorm.io/gorm"
)

// CategoryRepository ...
type CategoryRepository struct {
	*gorm.DB
}

// Fetch will fetch content from db
func (rp *CategoryRepository) Fetch(ctx context.Context) ([]domain.Category, error) {
	dao := []Category{}
	if err := rp.DB.WithContext(ctx).Find(&dao).Error; err != nil {
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}

	items := []domain.Category{}

	for i := range dao {
		r := convertCategoryToEntity(&dao[i])
		items = append(items, *r)
	}

	return items, nil
}

// Find will find content from db
func (rp *CategoryRepository) Find(ctx context.Context, id int) (*domain.Category, error) {
	dao := Category{}
	if err := rp.DB.WithContext(ctx).First(&dao, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound.Wrap(err)
		}
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}

	return convertCategoryToEntity(&dao), nil
}

// Store will create data to db
func (rp *CategoryRepository) Store(ctx context.Context, dm *domain.Category) error {
	dao := convertCategoryToDao(dm)
	if err := rp.DB.WithContext(ctx).Create(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	*dm = *convertCategoryToEntity(dao)

	return nil
}
