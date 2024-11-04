package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type MenuRepository interface {
	Create(ctx context.Context, req *requests.AddMenuRequest, imageURL string) error
	FindAll(ctx context.Context) ([]models.Menu, error)
	FindByID(ctx context.Context, email string) (*models.Menu, error)
	Edit(ctx context.Context, req *requests.EditMenuRequest, imageURL string) error
	Delete(ctx context.Context, id string) error
	FindByName(ctx context.Context, name string) (*models.Menu, error)
}
