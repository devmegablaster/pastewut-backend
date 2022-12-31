package middlewares

import (
  "github.com/gofiber/fiber/v2"

  "github.com/devmegablaster/pastewut-backend/pkg/errors"
  "github.com/devmegablaster/pastewut-backend/pkg/db"
  "github.com/devmegablaster/pastewut-backend/pkg/models"
)

// EditPastewutMiddleware is a middleware that checks if the pastewut exists and if the user is the owner of the pastewut
func EditPastewutMiddleware(c *fiber.Ctx) error {
  code := c.Params("code")
  if code == "" {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewutCode.Error())
  }

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

  var dbPastewut models.PasteWut

  if err := db.PsqlDB.Where("code = ?", code).First(&dbPastewut).Error; err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewutCode.Error())
  }

  if email != dbPastewut.Author {
    return c.Status(fiber.StatusUnauthorized).JSON(errors.Unauthorized.Error())
  }

  c.Locals("pastewut", dbPastewut)

  return c.Next()
}
