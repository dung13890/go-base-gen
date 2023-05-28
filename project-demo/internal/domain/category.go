//go:generate mockgen -source=$GOFILE -destination=mock/category_mock.go

package domain

import (
	"context"
	"time"
)

// Category entity
type Category struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// CategoryRepository represent the Category's repository contract
type CategoryRepository interface {
	Fetch(ctx context.Context) ([]Category, error)
	Find(ctx context.Context, id int) (*Category, error)
	Store(ctx context.Context, dm *Category) error
}

// CategoryUsecase represent the Category's usecase contract
type CategoryUsecase interface {
	Fetch(ctx context.Context) ([]Category, error)
	Find(ctx context.Context, id int) (*Category, error)
	Store(ctx context.Context, dm *Category) error
}
