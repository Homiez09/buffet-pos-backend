package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryGormRepository struct {
	DB *gorm.DB
}

func NewCategoryGormRepository(db *gorm.DB) *CategoryGormRepository {
	return &CategoryGormRepository{
		DB: db,
	}
}

func (t *CategoryGormRepository) Create(ctx context.Context, req *requests.AddCategoryRequest) error {
	// Generate UUID
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	category := &models.Category{
		ID:   id,
		Name: req.CategoryName,
	}

	result := t.DB.Create(category)

	return result.Error
}

func (t *CategoryGormRepository) FindAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	result := t.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (t *CategoryGormRepository) FindByID(ctx context.Context, categoryID string) (*models.Category, error) {
	var category models.Category
	result := t.DB.Where("id = ?", categoryID).First(&category)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (t *CategoryGormRepository) FindByName(ctx context.Context, categoryName string) (*models.Category, error) {
	var category models.Category
	result := t.DB.Where("name = ?", categoryName).First(&category)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
