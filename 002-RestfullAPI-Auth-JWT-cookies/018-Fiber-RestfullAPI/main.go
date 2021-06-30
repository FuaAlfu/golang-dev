package main

import (
	"book"

	"github.com/gofiber/fiber"

)

func main() {
	handleRequest()
}

//c will do the magic
func homePage(c *fiber.Ctx) {
	c.Send("Home Page :: welcome!")
}

func setupAllRoutes(app *fiber.App) {
	app.Get("/", homePage)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func handleRequest() {
	//initialize..
	app := fiber.New()

	setupAllRoutes(app)

	app.Listen(3000)
}
