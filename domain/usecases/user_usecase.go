package usecases

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
)

type UserUseCase interface {
	Register(ctx context.Context, req *requests.UserRegisterRequest) error
	Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserLoginResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
	config   *configs.Config
}

func NewUserService(userRepo repositories.UserRepository, config *configs.Config) UserUseCase {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}
