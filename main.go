package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Static("/", "./public")

	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":3000"))

}
