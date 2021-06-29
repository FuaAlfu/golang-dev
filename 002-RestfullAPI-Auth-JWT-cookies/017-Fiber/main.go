package main

import "github.com/gofiber/fiber"

func main() {
	handleRequest()
}

//c will do the magic
func homePage(c *fiber.Ctx) {
	c.Send("Home Page :: welcome!")
}

func handleRequest() {
	//initialize
	app := fiber.New()

	app.Get("/", homePage)

	app.Listen(3000)
}
