package response

import "time"

type RequestTableResponse struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"  validate:"required"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	TableID     uint      `json:"table_id"`
	CustomerID  uint      `json:"customer_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
