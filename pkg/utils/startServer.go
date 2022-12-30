package utils

import (
	"fmt"
  "os"

	"github.com/gofiber/fiber/v2"
)

func StartServer(app *fiber.App) {
  if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
    panic(err)
  }
}
