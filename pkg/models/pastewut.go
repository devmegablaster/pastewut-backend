package models

import (
  "math/rand"
  "time"
)

type PasteWut struct {
  Code string `json:"code" gorm:"unique"`
  Content string `json:"content" gorm:"not null"`
  Author string `json:"author,omitempty"`
}

func (pastewut *PasteWut) GenerateCode() {
  var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
  const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := make([]byte, 10)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}

  pastewut.Code = string(code)
}
