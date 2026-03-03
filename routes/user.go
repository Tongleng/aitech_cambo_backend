package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, ctrl *controllers.UserController) {
	users := router.Group("/users")

	users.Get("/", ctrl.GetUsers)
	users.Post("/register", ctrl.Register)
	users.Post("/login", ctrl.Login)
}
