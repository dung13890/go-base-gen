package http

import (
	"project-demo/internal/modules/ecomm3/usecase"

	"github.com/labstack/echo/v4"
)

// NewHandler will initialize the ecomm3 module endpoints
func NewHandler(g *echo.Group, uc *usecase.Usecase) {
	// categoryHandler will initialize the category/ resources endpoint
	categoryHandler := &CategoryHandler{
		Usecase: uc.CategoryUC,
	}

	// Category routes
	g.GET("/categorys", categoryHandler.Index)
	g.GET("/categorys/:id", categoryHandler.Show)
	g.POST("/categorys", categoryHandler.Store)
}
