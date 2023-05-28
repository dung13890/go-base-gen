package repository

import (
	"context"
	"time"

	"project-demo/internal/constants"
	"project-demo/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PasswordResetRepository ...
type PasswordResetRepository struct {
	*gorm.DB
}

// StoreOrUpdate will store or update password reset by email
func (rp *PasswordResetRepository) StoreOrUpdate(ctx context.Context, email, token string) error {
	dao := &PasswordReset{
		Email: email,
		Token: token,
	}

	if err := rp.DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"token", "created_at", "updated_at"}),
	}).Create(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	return nil
}

// FindEmailByToken will find password reset by token
func (rp *PasswordResetRepository) FindEmailByToken(ctx context.Context, token string) (string, error) {
	dao := &PasswordReset{
		Token: token,
	}

	createdAt := time.Now().Add(-constants.TokenResetPasswordLifetime)

	if err := rp.DB.WithContext(ctx).Where("created_at >= ?", createdAt).First(&dao, &dao).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.ErrAuthInvalidateToken.Wrap(err)
		}
		return "", errors.ErrUnexpectedDBError.Wrap(err)
	}

	return dao.Email, nil
}

// Delete will delete password reset by token
func (rp *PasswordResetRepository) Delete(ctx context.Context, email, token string) error {
	dao := &PasswordReset{
		Email: email,
		Token: token,
	}

	if err := rp.DB.WithContext(ctx).Delete(&dao, &dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	return nil
}
