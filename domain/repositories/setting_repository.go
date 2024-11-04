package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
)

type SettingRepository interface {
	GetSetting(ctx context.Context, key string) (*models.Setting, error)
	UpdateSetting(ctx context.Context, key string, value string) error
	AddSetting(ctx context.Context, key string, value string) error
}
