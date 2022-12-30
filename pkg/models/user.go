package models

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/devmegablaster/pastewut-backend/pkg/errors"
)

type User struct {
  Email string `json:"email" gorm:"primary_key"`
  Password string `json:"password" gorm:"not null"`
  Pastes []PasteWut `json:"pastes" gorm:"foreignkey:Code"`
}

func (u *User) ValidateEmail() error {
  if u.Email == "" {
    return errors.InvalidEmail.Err
  }

  emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
  if !emailRegex.MatchString(u.Email) {
    return errors.InvalidEmail.Err
  }

  return nil
}

func (u *User) ValidatePassword() error {
  if u.Password == "" {
    return errors.InvalidPassword.Err
  }

  passwordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`)
  if !passwordRegex.MatchString(u.Password) {
    return errors.InvalidPassword.Err
  }

  return nil
}

func (u *User) HashPassword() error {
  bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
  if err != nil {
    return errors.InternalServerError.Err
  }

  u.Password = string(bytes)

  return nil
}
