package usecase

import (
	"context"

	"project-demo/internal/domain"
	"project-demo/pkg/errors"
)

// ProductUsecase ...
type ProductUsecase struct {
	repo domain.ProductRepository
}

// NewProductUsecase will create new an productUsecase object representation of domain.ProductUsecase interface
func NewProductUsecase(repo domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repo: repo,
	}
}

// Fetch will fetch content from repo
func (uc *ProductUsecase) Fetch(ctx context.Context) ([]domain.Product, error) {
	items, err := uc.repo.Fetch(ctx)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return items, nil
}

// Find will find content from repo
func (uc *ProductUsecase) Find(ctx context.Context, id int) (*domain.Product, error) {
	item, err := uc.repo.Find(ctx, id)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *ProductUsecase) Store(ctx context.Context, dm *domain.Product) error {
	if err := uc.repo.Store(ctx, dm); err != nil {
		return errors.Throw(err)
	}

	return nil
}
