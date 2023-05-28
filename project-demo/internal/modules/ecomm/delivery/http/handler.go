package http

import (
	"project-demo/internal/modules/ecomm/usecase"

	"github.com/labstack/echo/v4"
)

// NewHandler will initialize the ecomm module endpoints
func NewHandler(g *echo.Group, uc *usecase.Usecase) {
	// projectHandler will initialize the project/ resources endpoint
	projectHandler := &ProjectHandler{
		Usecase: uc.ProjectUC,
	}

	// Project routes
	g.GET("/projects", projectHandler.Index)
	g.GET("/projects/:id", projectHandler.Show)
	g.POST("/projects", projectHandler.Store)
}
