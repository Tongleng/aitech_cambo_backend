package handlers

import (
	"backend/configs"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Image is required",
		})
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	url, err := configs.UploadFile(src, filename)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"url": url,
	})
}
