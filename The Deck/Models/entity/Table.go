package entity

import "time"

type Table struct {
	Id        uint      `json:"id"`
	Number    int       `json:"number" gorm:"type:int" validate:"required"`
	Capacity  int       `json:"capacity" gorm:"type:int" validate:"required"`
	Status    string    `json:"status" gorm:"type:varchar(20)" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
