package http

import (
	"time"

	"project-demo/internal/domain"
)

// CategoryRequest is request for create category
type CategoryRequest struct {
}

// CategoryResponse is struct used for category
type CategoryResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// convertCategoryRequestToEntity DTO
func convertCategoryRequestToEntity(request *CategoryRequest) *domain.Category {
	return &domain.Category{}
}

// convertCategoryEntityToResponse DTO
func convertCategoryEntityToResponse(dm *domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:        dm.ID,
		CreatedAt: dm.CreatedAt,
		UpdatedAt: dm.UpdatedAt,
	}
}
