package http

import (
	"project-demo/internal/modules/ecomm2/usecase"

	"github.com/labstack/echo/v4"
)

// NewHandler will initialize the ecomm2 module endpoints
func NewHandler(g *echo.Group, uc *usecase.Usecase) {
	// productHandler will initialize the product/ resources endpoint
	productHandler := &ProductHandler{
		Usecase: uc.ProductUC,
	}

	// Product routes
	g.GET("/products", productHandler.Index)
	g.GET("/products/:id", productHandler.Show)
	g.POST("/products", productHandler.Store)
}
