package http

import (
	"project-demo/config"
	"project-demo/internal/constants"
	"project-demo/internal/domain"
	"project-demo/internal/impl/service"
	"project-demo/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// SetupJWT .-
func SetupJWT() echo.MiddlewareFunc {
	jwtConf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(service.Claims)
		},
		SigningKey: []byte(config.GetAppConfig().AppJWTKey),
	}

	return echojwt.WithConfig(jwtConf)
}

// Authenticated .-
func Authenticated(svc domain.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user")
			user, err := svc.Decode(c.Request().Context(), token)
			if err != nil {
				return errors.Throw(err)
			}

			c.Set(constants.GuardJWT, user)

			return next(c)
		}
	}
}
