package entity

import "time"

type RequestTable struct {
	Id          uint      `json:"id"`
	Description string    `json:"description" gorm:"type:text" validate:"required"`
	Status      string    `json:"status" gorm:"type:varchar(20)" validate:"required"`
	StartDate   time.Time `json:"start_date" gorm:"type:varchar(100)" validate:"required"`
	EndDate     time.Time `json:"end_date" gorm:"type:varchar(100)" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	TableID     uint      `json:"table_id" gorm:"index"`
	CustomerID  uint      `json:"customer_id" gorm:"index"`
	Table       Table     `json:"table" gorm:"foreignKey:TableID"`
	Customer    Customer  `json:"customer" gorm:"foreignKey:CustomerID"`
}
