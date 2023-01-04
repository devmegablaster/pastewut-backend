package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/devmegablaster/pastewut-backend/pkg/db"
	"github.com/devmegablaster/pastewut-backend/pkg/errors"
	"github.com/devmegablaster/pastewut-backend/pkg/models"
)

func CreatePasteWut(c *fiber.Ctx) error {
  pastewut := new(models.PasteWut)
  if err := c.BodyParser(pastewut); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  if pastewut.Content == "" {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewut.Error())
  }

  pastewut.GenerateCode()

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
    return c.Status(fiber.StatusNotFound).JSON(errors.InvalidPastewutCode.Error())
  }

  return c.JSON(pastewut)
}

func CreateCustomPasteWut(c *fiber.Ctx) error {
  pastewut := new(models.PasteWut)
  if err := c.BodyParser(pastewut); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  if pastewut.Content == "" {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewut.Error())
  }

  if pastewut.Code == "" {
    pastewut.GenerateCode()
  }

  if err := db.PsqlDB.Create(&models.PasteWut{
    Code: pastewut.Code,
    Content: pastewut.Content,
    Author: c.Locals("email").(string),
  }).Error; err != nil {
    fmt.Println(err)
    return c.Status(fiber.StatusInternalServerError).JSON(errors.PastewutAlreadyExists.Error())
  }

  var pasteMap []models.PasteWut
  pasteMap = append(pasteMap, *pastewut)

  fmt.Println(pasteMap)

  return c.JSON(fiber.Map{
    "success": true,
    "code": pastewut.Code,
  })
}

func EditPasteWut(c *fiber.Ctx) error {
  pastewut := new(models.PasteWut)
  if err := c.BodyParser(pastewut); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  dbPastewut := c.Locals("pastewut").(models.PasteWut)

  if pastewut.Content == "" {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPastewut.Error())
  }

  if err := db.PsqlDB.Model(&models.PasteWut{}).Where("code = ?", dbPastewut.Code).Update("content", pastewut.Content).Error; err != nil {
    fmt.Println(err)
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  return c.JSON(fiber.Map{
    "success": true,
    "code": dbPastewut.Code,
    "content": pastewut.Content,
  })
}

func CheckPasteWutCode(c *fiber.Ctx) error {
  code := c.Params("code")
  if code == "" {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
      "exists": true,
    })
  }

  var pastewut models.PasteWut

  if err := db.PsqlDB.Where("code = ?", code).First(&pastewut).Error; err != nil {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
      "exists": true,
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "exists": false,
  })
}
