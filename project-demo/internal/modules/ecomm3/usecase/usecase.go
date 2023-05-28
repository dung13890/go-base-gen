package usecase

import (
	"project-demo/internal/domain"
	"project-demo/internal/modules/ecomm3/repository"
)

// Usecase for module auth
type Usecase struct {
	CategoryUC domain.CategoryUsecase
}

// NewUsecase implements from interface
func NewUsecase(
	repo *repository.Repository,
) *Usecase {
	return &Usecase{
		CategoryUC: NewCategoryUsecase(repo.CategoryR),
	}
}
