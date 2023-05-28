package http

import (
	"net/http"
	"strconv"

	"project-demo/internal/domain"
	"project-demo/pkg/errors"

	"github.com/labstack/echo/v4"
)

// ProductHandler represent the http handler
type ProductHandler struct {
	Usecase domain.ProductUsecase
}

// Index will fetch data
func (hl *ProductHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	items, err := hl.Usecase.Fetch(ctx)
	if err != nil {
		return errors.Throw(err)
	}
	itemsRes := make([]ProductResponse, 0)
	for i := range items {
		item := convertProductEntityToResponse(&items[i])
		itemsRes = append(itemsRes, item)
	}

	return c.JSON(http.StatusOK, itemsRes)
}

// Show will Find data
func (hl *ProductHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}
	ctx := c.Request().Context()
	item, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, convertProductEntityToResponse(item))
}

// Store will create data
func (hl *ProductHandler) Store(c echo.Context) error {
	request := new(ProductRequest)
	if err := c.Bind(request); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err := c.Validate(request)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	entity := convertProductRequestToEntity(request)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, entity); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, ProductStatusResponse{Status: true})
}
