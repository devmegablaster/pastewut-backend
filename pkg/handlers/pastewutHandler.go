package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/devmegablaster/pastewut-backend/pkg/db"
	"github.com/devmegablaster/pastewut-backend/pkg/helpers"
	"github.com/devmegablaster/pastewut-backend/pkg/models"
  "github.com/devmegablaster/pastewut-backend/pkg/errors"
)

func CreatePasteWut(c *fiber.Ctx) error {
  pastewut := new(models.PasteWut)
  if err := c.BodyParser(pastewut); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  if pastewut.Content == "" {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewut.Error())
  }

  pastewut.Code = helpers.GenerateCode()

  if err := db.PsqlDB.Create(&models.PasteWut{
    Code: pastewut.Code,
    Content: pastewut.Content,
  }).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  return c.JSON(fiber.Map{
    "success": true,
    "code": pastewut.Code,
  })
}

func GetPasteWut(c *fiber.Ctx) error {
  code := c.Params("code")
  if code == "" {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewutCode.Error())
  }

  var pastewut models.PasteWut
  if err := db.PsqlDB.Where("code = ?", code).First(&pastewut).Error; err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewutCode.Error())
  }

  return c.JSON(fiber.Map{
    "success": true,
    "content": pastewut.Content,
  })
}
