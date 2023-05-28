package config

import (
	"project-demo/pkg/errors"
	"project-demo/pkg/logger"

	"github.com/spf13/viper"
)

// AppConfig App Common
type AppConfig struct {
	App           string `mapstructure:"APP_ENV"`
	AllowedOrigin string `mapstructure:"APP_ALLOWED_ORIGIN"`
	AppHost       string `mapstructure:"APP_HOST"`
	AppJWTKey     string `mapstructure:"APP_JWT_KEY"`
	AppTimeZone   string `mapstructure:"APP_TIME_ZONE"`
}

// LoadConfig config setting from .env.
func LoadConfig() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	return nil
}

// GetAppConfig Unmarshal App Config from env
func GetAppConfig() AppConfig {
	c := AppConfig{}
	if err := viper.Unmarshal(&c); err != nil {
		logger.Error().Fatal(err)
	}

	return c
}
