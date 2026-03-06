package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func StoreRoutes(router fiber.Router, ctrl *controllers.StoreController) {
	storeRoutes := router.Group("/store")

	storeRoutes.Get("/", ctrl.GetAll)
	storeRoutes.Get("/:id", ctrl.GetByID)
	storeRoutes.Post("/", ctrl.Create)
	storeRoutes.Patch("/:id", ctrl.Update)
}
