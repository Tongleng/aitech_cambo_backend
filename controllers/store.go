package controllers

import (
	"backend/models"
	"backend/services"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StoreController struct {
	Service services.StoreService
}

func (c *StoreController) Create(ctx *fiber.Ctx) error {
	var store models.Store
	if err := ctx.BodyParser(&store); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := ctx.BodyParser(&store); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Service.CreateStore(&store); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(store)
}

func (c *StoreController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid store ID"})
	}

	store, err := c.Service.GetStoreByID(uint(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Store not found"})
	}

	return ctx.JSON(store)
}

func (c *StoreController) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	stores, total, err := c.Service.GetStores(page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"data":      stores,
		"total":     total,
		"page":      page,
		"last_page": (int(total) + limit - 1) / limit,
	})
}

func (c *StoreController) GetByCategory(ctx *fiber.Ctx) error {
	categoryID, err := ctx.ParamsInt("categoryId")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid Category ID"})
	}

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "12"))

	stores, total, err := c.Service.GetStoresByCategory(uint(categoryID), page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Reuse your pagination utility here
	return ctx.JSON(fiber.Map{
		"data":      stores,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total) / float64(limit)),
	})
}

func (c *StoreController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var store models.Store
	if err := ctx.BodyParser(&store); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON body"})
	}

	updatedStore, err := c.Service.UpdateStore(uint(id), &store)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to update store"})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Store updated successfully",
		"data":    updatedStore,
	})
}
