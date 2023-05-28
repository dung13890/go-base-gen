package http

import (
	"time"

	"project-demo/internal/domain"
)

// ProjectRequest is request for create project
type ProjectRequest struct {
}

// ProjectResponse is struct used for project
type ProjectResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// convertProjectRequestToEntity DTO
func convertProjectRequestToEntity(request *ProjectRequest) *domain.Project {
	return &domain.Project{}
}

// convertProjectEntityToResponse DTO
func convertProjectEntityToResponse(dm *domain.Project) ProjectResponse {
	return ProjectResponse{
		ID:        dm.ID,
		CreatedAt: dm.CreatedAt,
		UpdatedAt: dm.UpdatedAt,
	}
}
