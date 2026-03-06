package main

import (
	"backend/configs"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Use(logger.New())

	configs.ConnectDB()
	configs.RunMigrations()

	// configs.InitStorage()

	routes.Setup(app, configs.DB)

	app.Listen(":8080")
}
