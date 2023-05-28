package repository

import (
	"project-demo/internal/domain"

	"gorm.io/gorm"
)

// Repository for module auth
type Repository struct {
	CategoryR domain.CategoryRepository
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		CategoryR: &CategoryRepository{DB: db},
	}
}
