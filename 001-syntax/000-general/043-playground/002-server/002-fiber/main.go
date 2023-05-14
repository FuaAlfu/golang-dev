package main

import (
    "fmt"
    "net/http"

    "github.com/gofiber/fiber/v2"

)

func main() {
    // Create a new Fiber app.
    app := fiber.New()

    // Handle the "/" route.
    app.Get("/", func(c *fiber.Ctx) {
        c.Send("Hello, world!")
    })

    // Handle the "/about" route.
    app.Get("/about", func(c *fiber.Ctx) {
        c.Send("This is the about page.")
    })

    // Listen on port 8080.
    app.Listen(":8080")
}