package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductCategoryRoutes(router fiber.Router, ctrl *controllers.ProductCategoryController) {
	productCategoryRoutes := router.Group("/product-categories")

	productCategoryRoutes.Post("/", ctrl.CreateProductCategory)
	productCategoryRoutes.Put("/:id", ctrl.UpdateProductCategory)
	productCategoryRoutes.Get("/", ctrl.GetProductCategory)
}
