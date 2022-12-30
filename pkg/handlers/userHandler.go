package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/devmegablaster/pastewut-backend/pkg/db"
	"github.com/devmegablaster/pastewut-backend/pkg/errors"
	"github.com/devmegablaster/pastewut-backend/pkg/models"
)

func RegisterUser(c *fiber.Ctx) error {
  user := new(models.User)
  if err := c.BodyParser(user); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidDetails.Error())
  }

  if err := user.ValidateEmail(); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidEmail.Error())
  }

  if err := user.HashPassword(); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  if err := db.PsqlDB.Create(&models.User{
    Email: user.Email,
    Password: user.Password,
  }).Error; err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.UserAlreadyExists.Error())
  }

  return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    "message": "User created successfully",
  })
}

func Login(c *fiber.Ctx) error {
  user := new(models.User)
  if err := c.BodyParser(user); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidDetails.Error())
  }

  if err := user.ValidateEmail(); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidEmail.Error())
  }

  dbUser := new(models.User)

  if err := db.PsqlDB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.UserNotFound.Error())
  }

  if err := dbUser.ComparePassword(user.Password); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(errors.InvalidPassword.Error())
  }

  token, err := user.GenerateJWT()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(errors.InternalServerError.Error())
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "User logged in successfully",
    "token": token,
  })
}
