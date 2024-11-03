package usecases

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type CategoryUseCase interface {
	AddCategory(ctx context.Context, req *requests.AddCategoryRequest) error
	FindAllCategories(ctx context.Context) ([]responses.CategoryDetail, error)
	FindCategoryByID(ctx context.Context, categoryID string) (*responses.CategoryDetail, error)
	DeleteCategory(ctx context.Context, categoryID string) error
}

type CategoryService struct {
	categoryRepo repositories.CategoryRepository
	config       *configs.Config
}

func NewCategoryService(categoryRepo repositories.CategoryRepository, config *configs.Config) CategoryUseCase {
	return &CategoryService{
		categoryRepo: categoryRepo,
		config:       config,
	}
}

func (c *CategoryService) AddCategory(ctx context.Context, req *requests.AddCategoryRequest) error {
	category, err := c.categoryRepo.FindByName(ctx, req.CategoryName)
	if err != nil {
		return err
	}

	if category != nil {
		return exceptions.ErrDuplicatedCategoryName
	}
	return c.categoryRepo.Create(ctx, req)
}

func (c *CategoryService) FindAllCategories(ctx context.Context) ([]responses.CategoryDetail, error) {
	categories, err := c.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	categoryDetails := make([]responses.CategoryDetail, 0)
	for _, category := range categories {
		categoryDetails = append(categoryDetails, responses.CategoryDetail{
			BaseCategory: responses.BaseCategory{
				ID:        category.ID,
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			},
		})
	}
	return categoryDetails, nil
}

func (c *CategoryService) FindCategoryByID(ctx context.Context, categoryID string) (*responses.CategoryDetail, error) {
	category, err := c.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, exceptions.ErrCategoryNotFound
	}

	return &responses.CategoryDetail{
		BaseCategory: responses.BaseCategory{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
	}, nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, categoryID string) error {
	category, err := c.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return err
	}
	if category == nil {
		return exceptions.ErrCategoryNotFound
	}
	return c.categoryRepo.Delete(ctx, categoryID)
}
