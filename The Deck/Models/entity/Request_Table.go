package entity

import "time"

type RequestTable struct {
	Id         uint      `json:"id"`
	TableID    uint      `json:"table_id"`
	Table      Table     `json:"tables" gorm:"foreignKey:table_id"`
	CustomerID uint      `json:"customer_id"`
	Customer   Customer  `json:"customers" gorm:"foreignKey:customer_id"`
	AdminID    *uint     `json:"admin_id"`
	Admin      Admin     `json:"admins" gorm:"foreignKey:admin_id"`
	Status     int       `json:"status"`
	Notes      string    `json:"note"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
