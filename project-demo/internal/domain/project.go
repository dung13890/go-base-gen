//go:generate mockgen -source=$GOFILE -destination=mock/project_mock.go

package domain

import (
	"context"
	"time"
)

// Project entity
type Project struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// ProjectRepository represent the Project's repository contract
type ProjectRepository interface {
	Fetch(ctx context.Context) ([]Project, error)
	Find(ctx context.Context, id int) (*Project, error)
	Store(ctx context.Context, dm *Project) error
}

// ProjectUsecase represent the Project's usecase contract
type ProjectUsecase interface {
	Fetch(ctx context.Context) ([]Project, error)
	Find(ctx context.Context, id int) (*Project, error)
	Store(ctx context.Context, dm *Project) error
}
