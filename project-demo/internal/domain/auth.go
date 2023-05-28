//go:generate mockgen -source=$GOFILE -destination=mock/auth_mock.go

package domain

import (
	"context"
)

// AuthUsecase represent the auth's usecase contract
type AuthUsecase interface {
	// Register new user
	Register(ctx context.Context, u *User) (*User, error)
	// Login user
	Login(ctx context.Context, u *User, ip string) (string, int64, error)
	// Logout user
	Logout(ctx context.Context, token any) error
}
