package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type TableRepository interface {
	Create(ctx context.Context, req *requests.AddTableRequest) error
	FindByID(ctx context.Context, email string) (*models.Table, error)
	FindByName(ctx context.Context, name string) (*models.Table, error)
}
