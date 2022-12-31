package db

import (
	"fmt"
  "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

  "github.com/devmegablaster/pastewut-backend/pkg/models"
)

var PsqlDB *gorm.DB

func InitiatePostgres() {
  var err error

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
  PsqlDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  fmt.Println("Postgres connected")

  // Migrate the schema
  PsqlDB.AutoMigrate(&models.User{})
  PsqlDB.AutoMigrate(&models.PasteWut{})
}
