package usecase

import (
	"context"

	"project-demo/internal/domain"
	"project-demo/pkg/errors"
)

// CategoryUsecase ...
type CategoryUsecase struct {
	repo domain.CategoryRepository
}

// NewCategoryUsecase will create new an categoryUsecase object representation of domain.CategoryUsecase interface
func NewCategoryUsecase(repo domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
	}
}

// Fetch will fetch content from repo
func (uc *CategoryUsecase) Fetch(ctx context.Context) ([]domain.Category, error) {
	items, err := uc.repo.Fetch(ctx)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return items, nil
}

// Find will find content from repo
func (uc *CategoryUsecase) Find(ctx context.Context, id int) (*domain.Category, error) {
	item, err := uc.repo.Find(ctx, id)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *CategoryUsecase) Store(ctx context.Context, dm *domain.Category) error {
	if err := uc.repo.Store(ctx, dm); err != nil {
		return errors.Throw(err)
	}

	return nil
}
