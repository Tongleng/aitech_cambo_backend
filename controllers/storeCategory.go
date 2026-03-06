package controllers

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

type StoreCategoryController struct {
	Service services.StoreCategoryService
}

func (c *StoreCategoryController) CreateStoreCategory(ctx *fiber.Ctx) error {
	var category models.StoreCategory

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	res, err := c.Service.CreateStoreCategory(&category)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(res)
}

func (c *StoreCategoryController) GetStoreCategory(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	res, err := c.Service.GetStoreCategoryPaginated(page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}

func (c *StoreCategoryController) UpdateStoreCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var category models.StoreCategory
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := c.Service.UpdateStoreCategory(uint(id), &category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"message": "Store category updated successfully",
		"id":      id,
	})
}
