package http

import (
	"net/http"
	"strconv"

	"project-demo/internal/domain"
	"project-demo/pkg/errors"

	"github.com/labstack/echo/v4"
)

// CategoryHandler represent the http handler
type CategoryHandler struct {
	Usecase domain.CategoryUsecase
}

// Index will fetch data
func (hl *CategoryHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	items, err := hl.Usecase.Fetch(ctx)
	if err != nil {
		return errors.Throw(err)
	}
	itemsRes := make([]CategoryResponse, 0)
	for i := range items {
		item := convertCategoryEntityToResponse(&items[i])
		itemsRes = append(itemsRes, item)
	}

	return c.JSON(http.StatusOK, itemsRes)
}

// Show will Find data
func (hl *CategoryHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}
	ctx := c.Request().Context()
	item, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, convertCategoryEntityToResponse(item))
}

// Store will create data
func (hl *CategoryHandler) Store(c echo.Context) error {
	request := new(CategoryRequest)
	if err := c.Bind(request); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err := c.Validate(request)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	entity := convertCategoryRequestToEntity(request)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, entity); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, CategoryStatusResponse{Status: true})
}
