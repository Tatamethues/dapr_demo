package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)

	// Start server
	log.Fatal(app.Listen(":80"))
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hello world from fiber!",
	})
}
