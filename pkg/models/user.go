package models

type User struct {
  ID uint `json:"id" gorm:"primary_key"`
  Email string `json:"email" gorm:"unique"`
  Password string `json:"password" gorm:"not null"`
  Pastes []PasteWut `json:"pastes" gorm:"foreignkey:Code"`
}
