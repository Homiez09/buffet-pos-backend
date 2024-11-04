package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"gorm.io/gorm"
)

type SettingGormRepository struct {
	DB *gorm.DB
}

func NewSettingGormRepository(db *gorm.DB) *SettingGormRepository {
	return &SettingGormRepository{
		DB: db,
	}
}

func (s *SettingGormRepository) GetSetting(ctx context.Context, key string) (*models.Setting, error) {
	var setting models.Setting
	result := s.DB.Where("key = ?", key).First(&setting)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &setting, nil
}

func (s *SettingGormRepository) UpdateSetting(ctx context.Context, key string, value string) error {
	result := s.DB.Model(&models.Setting{}).Where("key = ?", key).Update("value", value)
	return result.Error
}

func (s *SettingGormRepository) AddSetting(ctx context.Context, key string, value string) error {
	setting := &models.Setting{
		Key:   key,
		Value: value,
	}
	result := s.DB.Create(setting)
	return result.Error
}
