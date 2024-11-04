package repositories

import "context"

type SettingRepository interface {
	GetSetting(ctx context.Context, key string) (string, error)
	UpdateSetting(ctx context.Context, key string, value string) error
	AddSetting(ctx context.Context, key string, value string) error
}
