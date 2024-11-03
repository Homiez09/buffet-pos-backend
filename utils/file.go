package utils

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func OpenFile(c *fiber.Ctx) (multipart.File, error) {
	// Pull form file
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get file",
		})
	}

	// Open File
	file, err := fileHeader.Open()
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to open file",
		})
	}
	defer file.Close()

	return file, nil
}
