package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type CategoryRepository interface {
	Create(ctx context.Context, req *requests.AddCategoryRequest) error
	FindAll(ctx context.Context) ([]models.Category, error)
	FindByID(ctx context.Context, id string) (*models.Category, error)
	FindByName(ctx context.Context, name string) (*models.Category, error)
	Delete(ctx context.Context, id string) error
}
