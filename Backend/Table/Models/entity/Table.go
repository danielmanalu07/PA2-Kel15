package entity

import "time"

type Table struct {
	Id        uint      `json:"id"`
	Number    int32     `json:"number" gorm:"type:int" validate:"required"`
	Capacity  int32     `json:"capacity" gorm:"type:int" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
