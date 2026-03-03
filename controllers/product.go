package controllers

import (
	"backend/models"
	"backend/services"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	Service services.ProductService
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Service.CreateProduct(&product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(product)
}

func (c *ProductController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	product, err := c.Service.GetProductByID(uint(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	return ctx.JSON(product)
}

func (c *ProductController) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	products, total, err := c.Service.GetProducts(page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": (int(total) + limit - 1) / limit,
	})
}

func (c *ProductController) GetByCategory(ctx *fiber.Ctx) error {
	categoryID, err := ctx.ParamsInt("categoryId")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid Category ID"})
	}

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "12"))

	products, total, err := c.Service.GetProductsByCategory(uint(categoryID), page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Reuse your pagination utility here
	return ctx.JSON(fiber.Map{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total) / float64(limit)),
	})
}
