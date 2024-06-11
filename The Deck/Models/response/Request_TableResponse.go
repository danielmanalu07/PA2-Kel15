package response

import "time"

type ResponseRequestTable struct {
	Id         uint      `json:"id"`
	TableID    uint      `json:"table_id"`
	CustomerID uint      `json:"customer_id"`
	Status     int       `json:"status"`
	Notes      string    `json:"notes"`
	AdminID    *uint     `json:"admin_id"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
