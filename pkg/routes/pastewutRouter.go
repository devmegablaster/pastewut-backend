package routes

import (
	"github.com/devmegablaster/pastewut-backend/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

func PasteWutRouter(app *fiber.App) {
  pastewut := app.Group("/pastewut")
  pastewut.Post("/", handlers.CreatePasteWut)
  pastewut.Get("/:code", handlers.GetPasteWut)
}
