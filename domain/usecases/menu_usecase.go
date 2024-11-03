package usecases

import (
	"context"
	"mime/multipart"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/storage"
)

type MenuUseCase interface {
	Create(ctx context.Context, req *requests.AddMenuRequest, file multipart.File) error
	FindAll(ctx context.Context) ([]responses.MenuDetail, error)
	FindByID(ctx context.Context, email string) (*responses.MenuDetail, error)
	DeleteMenu(ctx context.Context, id string) error
}

type menuService struct {
	menuRepo       repositories.MenuRepository
	config         *configs.Config
	storageService storage.StorageService
}

func NewMenuService(menuRepo repositories.MenuRepository, config *configs.Config, storageService storage.StorageService) MenuUseCase {
	return &menuService{
		menuRepo:       menuRepo,
		config:         config,
		storageService: storageService,
	}
}

func (m *menuService) Create(ctx context.Context, req *requests.AddMenuRequest, file multipart.File) error {
	imageURL, err := m.storageService.UploadFile(ctx, file)
	if err != nil {
		return err
	}
	return m.menuRepo.Create(ctx, req, imageURL)
}

func (m *menuService) FindAll(ctx context.Context) ([]responses.MenuDetail, error) {
	menus, err := m.menuRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	menuDetails := make([]responses.MenuDetail, 0)
	for _, menu := range menus {
		menuDetails = append(menuDetails, responses.MenuDetail{
			BaseMenu: responses.BaseMenu{
				ID:          menu.ID,
				Name:        menu.Name,
				Description: *menu.Description,
				CategoryID:  *menu.CategoryID,
				ImageURL:    *menu.ImageURL,
				IsAvailable: menu.IsAvailable,
			},
		})
	}
	return menuDetails, nil
}

func (m *menuService) FindByID(ctx context.Context, email string) (*responses.MenuDetail, error) {
	menu, err := m.menuRepo.FindByID(ctx, email)
	if err != nil {
		return nil, err
	}

	if menu == nil {
		return nil, exceptions.ErrMenuNotFound
	}

	return &responses.MenuDetail{
		BaseMenu: responses.BaseMenu{
			ID:          menu.ID,
			Name:        menu.Name,
			Description: *menu.Description,
			CategoryID:  *menu.CategoryID,
			ImageURL:    *menu.ImageURL,
			IsAvailable: menu.IsAvailable,
		},
	}, nil
}

func (m *menuService) DeleteMenu(ctx context.Context, id string) error {
	menu, err := m.menuRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if menu == nil {
		return exceptions.ErrMenuNotFound
	}

	return m.menuRepo.Delete(ctx, id)
}
