package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-microservice/config"
	"go-microservice/routes"
)

func main() {
	app := fiber.New()

	// Koneksikan ke database
	config.ConnectDatabase()

	// Add logger middleware to log all requests
	app.Use(logger.New(logger.Config{
		Format: "${time} [${ip}:${port}] ${status} - ${method} ${path}\n",
	}))

	// Setup rute
	routes.BookRoutes(app)

	app.Listen(":8080")
}
