package routes

import (
  "github.com/gofiber/fiber/v2"

  "github.com/devmegablaster/pastewut-backend/pkg/handlers"
)

func UserRouter(app *fiber.App) {
  user := app.Group("/user")
  user.Post("register", handlers.RegisterUser)
  user.Post("login", handlers.Login)
}
