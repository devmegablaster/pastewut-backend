package routes

import (
	"github.com/devmegablaster/pastewut-backend/pkg/handlers"
	"github.com/devmegablaster/pastewut-backend/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PasteWutRouter(app *fiber.App) {
  pastewut := app.Group("/pastewut")
  pastewut.Post("/", handlers.CreatePasteWut)
  pastewut.Get("/:code", handlers.GetPasteWut)
  pastewut.Post("/custom", middlewares.AuthMiddleware, handlers.CreateCustomPasteWut)
}
