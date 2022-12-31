package middlewares

import (
	"github.com/devmegablaster/pastewut-backend/pkg/db"
	"github.com/devmegablaster/pastewut-backend/pkg/errors"
	"github.com/devmegablaster/pastewut-backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware(c *fiber.Ctx) error {
  token := c.Get("Authorization")
  if token == "" {
    return c.Status(fiber.StatusUnauthorized).JSON(errors.Unauthorized.Error())
  }

  dbUser := new(models.User)
  email, err := dbUser.ValidateJWT(token)

  if err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(errors.Unauthorized.Error())
  }

  if err := db.PsqlDB.Where("email = ?", email).First(&dbUser).Error; err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(errors.Unauthorized.Error())
  }

  c.Locals("email", email)
  c.Locals("user", dbUser)

  return c.Next()
}
