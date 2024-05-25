package dto

type RequestTableCreate struct {
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	TableID     uint   `json:"table_id"`
}

type RequestTableUpdate struct {
	Status      string `json:"status" validate:"required"`
	Description string `json:"description" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	TableID     uint   `json:"table_id"`
}
