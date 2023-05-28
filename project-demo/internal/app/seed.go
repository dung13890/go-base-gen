package app

import (
	"context"

	"project-demo/config"
	"project-demo/internal/domain"
	authModule "project-demo/internal/modules/auth/repository"
	"project-demo/pkg/errors"
	"project-demo/pkg/postgres"

	"github.com/spf13/viper"
)

var pathJSON = "db/seeds/data.json"

type seedData struct {
	Roles []domain.Role `json:"roles"`
	Users []domain.User `json:"users"`
}

// Seed is function that seed data
func Seed(dbConfig config.Database) error {
	db, err := postgres.NewGormDB(dbConfig)
	if err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	authModule := authModule.NewRepository(db)

	viper.SetConfigFile(pathJSON)
	if err = viper.ReadInConfig(); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	data := seedData{}
	if err := viper.Unmarshal(&data); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	if err := data.seedAuth(authModule); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// seedAuth is function that seed auth data
func (seed *seedData) seedAuth(md *authModule.Repository) error {
	for i := range seed.Roles {
		if err := md.RoleR.Store(context.Background(), &seed.Roles[i]); err != nil {
			return errors.ErrInternalServerError.Wrap(err)
		}
	}

	for j := range seed.Users {
		if err := md.UserR.Store(context.Background(), &seed.Users[j]); err != nil {
			return errors.ErrInternalServerError.Wrap(err)
		}
	}

	return nil
}
