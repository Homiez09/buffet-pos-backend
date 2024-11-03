package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type TableRepository interface {
	Create(ctx context.Context, req *requests.AddTableRequest) error
	Edit(ctx context.Context, req *requests.EditTableRequest) error
	FindAll(ctx context.Context) ([]models.Table, error)
	FindByID(ctx context.Context, email string) (*models.Table, error)
	FindByName(ctx context.Context, name string) (*models.Table, error)
	Delete(ctx context.Context, id string) error
}
