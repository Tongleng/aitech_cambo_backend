package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router, ctrl *controllers.ProductController) {
	productRoutes := router.Group("/product")

	productRoutes.Post("/", ctrl.Create)
	productRoutes.Get("/", ctrl.GetAll)
	productRoutes.Get("/:id", ctrl.GetByID)
	productRoutes.Patch("/:id", ctrl.Update)
	productRoutes.Get("/category/:categoryId", ctrl.GetByCategory)
}
