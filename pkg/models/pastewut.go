package models

type PasteWut struct {
  Code string `json:"code" gorm:"unique"`
  Content string `json:"content" gorm:"not null"`
}
