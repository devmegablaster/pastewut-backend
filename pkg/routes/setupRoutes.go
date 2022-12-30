package routes

import (
  "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

  // Hello World
  app.Get("/", func (c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  // Setup routes here
  PasteWutRouter(app)
}
