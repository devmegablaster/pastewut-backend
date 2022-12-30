package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/devmegablaster/pastewut-backend/pkg/config"
	"github.com/devmegablaster/pastewut-backend/pkg/db"
	"github.com/devmegablaster/pastewut-backend/pkg/routes"
	"github.com/devmegablaster/pastewut-backend/pkg/utils"
)

func main() {
  config.LoadEnv()

  db.InitiatePostgres()

  app := fiber.New()
  routes.SetupRoutes(app)
  utils.StartServer(app)
}
