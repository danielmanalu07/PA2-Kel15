package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Cashier struct {
	Id         uint       `json:"id"`
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required, min:6" gorm:"not null"`
	Phone      string     `json:"phone" validate:"required, min:11" gorm:"not null"`
	Photo      string     `json:"photo" gorm:"null"`
	AdminID    uint       `gorm:"" json:"admin_id"`
	Admin      Admin      `gorm:"foreignKey:AdminID" json:"admin"`
	Created_at *time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at *time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

func (C *Cashier) ValidateCashier() error {
	validate := validator.New()
	return validate.Struct(C)
}
