package cloudinary

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/storage"
)

type CloudinaryStorageService struct {
	client *cloudinary.Cloudinary
	cfg    *configs.Config
}

func NewCloudinaryStorageService(cfg *configs.Config) storage.StorageService {
	cld, err := cloudinary.NewFromParams(cfg.CloudinaryCloudName, cfg.CloudinaryApiKey, cfg.CloudinaryApiSecret)
	if err != nil {
		log.Fatalf("Failed to create Cloudinary client: %v", err)
	}
	return &CloudinaryStorageService{client: cld}
}

func (c *CloudinaryStorageService) UploadFile(ctx context.Context, file multipart.File) (string, error) {
	uploadResult, err := c.client.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "buffet-pos",
	})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
