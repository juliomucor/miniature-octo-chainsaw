package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/juliomucor/miniature-octo-chainsaw/controllers"
	"log"
)

func main() {
	app := fiber.New()
	// using logger middleware, connecting it with the app
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

func setupRoutes(app *fiber.App) {
	//	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	// api group
	api := app.Group("/api")

	// given response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	controllers.TaskRoute(api.Group("/tasks"))
}
