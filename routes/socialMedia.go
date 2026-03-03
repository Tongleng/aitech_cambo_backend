package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SocialMediaRoutes(router fiber.Router, ctrl *controllers.SocialMediaController) {
	socialMedial := router.Group("/social-media")

	socialMedial.Get("/", ctrl.GetAll)
	socialMedial.Post("/", ctrl.Create)
	socialMedial.Put("/:id", ctrl.Update)
}
