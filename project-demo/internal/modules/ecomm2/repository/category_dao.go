package repository

import (
	"project-demo/internal/domain"
	"project-demo/pkg/utils"

	"gorm.io/gorm"
)

// Category DAO model
type Category struct {
	gorm.Model
}

// convertCategoryToEntity .-
func convertCategoryToEntity(dao *Category) *domain.Category {
	e := &domain.Category{
		ID:        dao.ID,
		CreatedAt: dao.CreatedAt,
		UpdatedAt: dao.UpdatedAt,
	}

	return e
}

// convertCategoryToDao .-
func convertCategoryToDao(entity *domain.Category) *Category {
	d := &Category{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
	}

	return d
}
