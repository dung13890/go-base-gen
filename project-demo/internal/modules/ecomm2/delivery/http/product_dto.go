package http

import (
	"time"

	"project-demo/internal/domain"
)

// ProductRequest is request for create product
type ProductRequest struct {
}

// ProductResponse is struct used for product
type ProductResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// convertProductRequestToEntity DTO
func convertProductRequestToEntity(request *ProductRequest) *domain.Product {
	return &domain.Product{}
}

// convertProductEntityToResponse DTO
func convertProductEntityToResponse(dm *domain.Product) ProductResponse {
	return ProductResponse{
		ID:        dm.ID,
		CreatedAt: dm.CreatedAt,
		UpdatedAt: dm.UpdatedAt,
	}
}
