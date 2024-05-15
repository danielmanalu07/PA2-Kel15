package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Admin struct {
	Id         uint      `json:"id"`
	Username   string    `json:"username" validate:"required" gorm:"not null"`
	Password   string    `json:"password" gorm:"not null" validate:"required"`
	Created_at time.Time `gorm:"autoCreateTime" json:"created_at"`
	Update_at  time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

func (A *Admin) ValidateAdmin() error {
	validate := validator.New()
	return validate.Struct(A)
}
