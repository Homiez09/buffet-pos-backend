package utils

import (
	"context"

	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

func OpenFile(c *fiber.Ctx) (multipart.File, error) {

	// Pull form file
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to get file",
			"status":  "400",
			"message": err.Error(),
		})
	}

	// Open File
	file, err := fileHeader.Open()
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to open file",
			"status":  "400",
			"message": err.Error(),
		})
	}
	defer file.Close()

	return file, nil
}

func UploadFile(ctx context.Context, cld *cloudinary.Cloudinary, file interface{}) (*uploader.UploadResult, error) {
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return nil, err
	}

	return res, err
}
