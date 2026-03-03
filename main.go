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

	// 2. Add CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:5173, http://localhost:3001",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	// 1. Database
	configs.ConnectDB()
	configs.RunMigrations()

	// 2. Setup Routes
	routes.Setup(app, configs.DB)

	// 3. Start
	app.Listen(":8080")
}
