package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-microservice/controller"
)

// BookRoutes mendefinisikan rute terkait buku
func BookRoutes(app *fiber.App) {
	app.Get("/books", controller.GetBooks)
	app.Post("/books", controller.CreateBook)
	app.Put("/books/:books_id", controller.UpdateBook)
	app.Delete("/books/:books_id", controller.DeleteBook)
}
