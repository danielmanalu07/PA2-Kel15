package dto

type RequestRequestTableCreate struct {
	Number   int    `json:"number" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

type RequestRequestTableUpdate struct {
	Number   int    `json:"number" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
	Status   string `json:"status" validate:"required"`
}
