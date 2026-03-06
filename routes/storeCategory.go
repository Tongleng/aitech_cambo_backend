package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func StoreCategoryRoutes(router fiber.Router, ctrl *controllers.StoreCategoryController) {
	storeCategoryRoutes := router.Group("/store-categories")

	storeCategoryRoutes.Post("/", ctrl.CreateStoreCategory)
	storeCategoryRoutes.Put("/:id", ctrl.UpdateStoreCategory)
	storeCategoryRoutes.Get("/", ctrl.GetStoreCategory)
}
