package repostories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
