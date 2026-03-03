package main

import (
	"backend/configs"
	"backend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	mainDNS := os.Getenv("MAIN_DNS")

	app.Use(cors.New(cors.Config{
		AllowOrigins:     mainDNS + ", http://localhost:3000, http://localhost:5173, http://localhost:3001",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	configs.ConnectDB()
	configs.RunMigrations()

	routes.Setup(app, configs.DB)

	app.Listen(":8080")
}
