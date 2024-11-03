package usecases

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
)

type MenuUseCase interface{}

type menuService struct {
	menuRepo repositories.MenuRepository
	config   *configs.Config
}

func NewMenuService(menuRepo repositories.MenuRepository, config *configs.Config) MenuUseCase {
	return &menuService{
		menuRepo: menuRepo,
		config:   config,
	}
}
