package controllers

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

type SocialMediaController struct {
	Service services.SocialMediaService
}

func (c *SocialMediaController) Create(ctx *fiber.Ctx) error {
	var social models.SocialMedia
	if err := ctx.BodyParser(&social); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := c.Service.Create(&social); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(social)
}

func (c *SocialMediaController) GetAll(ctx *fiber.Ctx) error {
	socials, err := c.Service.GetAllSocials()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(socials)
}

func (c *SocialMediaController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var social models.SocialMedia
	if err := ctx.BodyParser(&social); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Service.UpdateSocial(uint(id), &social); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Social media updated successfully"})
}
