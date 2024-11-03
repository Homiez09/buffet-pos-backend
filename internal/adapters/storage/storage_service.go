package storage

import (
	"context"
	"mime/multipart"
)

type StorageService interface {
	UploadFile(ctx context.Context, file multipart.File) (string, error)
}
