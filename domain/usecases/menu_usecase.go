package usecases

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type MenuUseCase interface {
	Create(ctx context.Context, req *requests.AddMenuRequest) error
	FindAll(ctx context.Context) ([]models.Menu, error)
	FindByID(ctx context.Context, email string) (*models.Menu, error)
}

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

func (m *menuService) Create(ctx context.Context, req *requests.AddMenuRequest) error {

	err := m.menuRepo.Create(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (m *menuService) FindAll(ctx context.Context) ([]models.Menu, error) {
	var res []models.Menu

	res, err := m.menuRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (m *menuService) FindByID(ctx context.Context, email string) (*models.Menu, error) {

	res, err := m.menuRepo.FindByID(ctx, email)
	if err != nil {
		return nil, err
	}
	return res, nil

}
