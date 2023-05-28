package usecase

import (
	"project-demo/internal/domain"
	"project-demo/internal/modules/ecomm2/repository"
)

// Usecase for module auth
type Usecase struct {
	ProductUC domain.ProductUsecase
}

// NewUsecase implements from interface
func NewUsecase(
	repo *repository.Repository,
) *Usecase {
	return &Usecase{
		ProductUC: NewProductUsecase(repo.ProductR),
	}
}
