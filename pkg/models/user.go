package models

import (
	"os"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/devmegablaster/pastewut-backend/pkg/errors"
	"github.com/golang-jwt/jwt"
)

type User struct {
  Email string `json:"email" gorm:"primary_key"`
  Password string `json:"password,omitempty" gorm:"not null"`
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

func (u *User) ComparePassword(password string) error {
  err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
  if err != nil {
    return errors.InvalidPassword.Err
  }

  return nil
}

func (u *User) GenerateJWT() (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)

  var claims *jwt.MapClaims

  claims = &jwt.MapClaims{
    "email": u.Email,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  }

  token.Claims = claims
  tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

  if err != nil {
    return "", errors.InternalServerError.Err
  }

  return tokenString, nil
}

func (u *User) ValidateJWT(tokenString string) (string, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.InvalidToken.Err
    }

    return []byte(os.Getenv("JWT_SECRET")), nil
  })

  if err != nil {
    return "", errors.InvalidToken.Err
  }

  if !token.Valid {
    return "", errors.InvalidToken.Err
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    return "", errors.InvalidToken.Err
  }

  email, ok := claims["email"].(string)
  if !ok {
    return "", errors.InvalidToken.Err
  }

  return email, nil
}
