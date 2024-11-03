package usecases

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
)

type MenuUseCase interface{}

type menuService struct {
	menuRepo   repositories.MenuRepository
	config     *configs.Config
	cloudinary *cloudinary.Cloudinary
}

func NewMenuService(menuRepo repositories.MenuRepository, config *configs.Config, cloudinary *cloudinary.Cloudinary) MenuUseCase {
	return &menuService{
		menuRepo:   menuRepo,
		config:     config,
		cloudinary: cloudinary,
	}
}
