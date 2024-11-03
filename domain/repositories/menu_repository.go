package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type MenuRepository interface {
	Create(ctx context.Context, req *requests.AddMenuRequest) error
	FindAll(ctx context.Context) ([]models.Menu, error)
	FindByID(ctx context.Context, email string) (*models.Menu, error)
}
