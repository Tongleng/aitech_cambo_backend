package controllers

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

type ProductCategoryController struct {
	Service services.ProductCategoryService
}

func (c *ProductCategoryController) CreateProductCategory(ctx *fiber.Ctx) error {
	var category models.ProductCategory

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	res, err := c.Service.CreateProductCategory(&category)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(res)
}

func (c *ProductCategoryController) GetProductCategory(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	res, err := c.Service.GetProductCategoryPaginated(page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(res)
}

func (c *ProductCategoryController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var product models.ProductCategory
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Service.UpdateProductCategory(uint(id), &product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Social media updated successfully"})
}

func (c *ProductCategoryController) UpdateProductCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var category models.ProductCategory
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := c.Service.UpdateProductCategory(uint(id), &category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"message": "Product category updated successfully",
		"id":      id,
	})
}
