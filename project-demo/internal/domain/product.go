//go:generate mockgen -source=$GOFILE -destination=mock/product_mock.go

package domain

import (
	"context"
	"time"
)

// Product entity
type Product struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// ProductRepository represent the Product's repository contract
type ProductRepository interface {
	Fetch(ctx context.Context) ([]Product, error)
	Find(ctx context.Context, id int) (*Product, error)
	Store(ctx context.Context, dm *Product) error
}

// ProductUsecase represent the Product's usecase contract
type ProductUsecase interface {
	Fetch(ctx context.Context) ([]Product, error)
	Find(ctx context.Context, id int) (*Product, error)
	Store(ctx context.Context, dm *Product) error
}
