package dto

type RequestTableCreate struct {
	Number   int `json:"number"  validate:"required"`
	Capacity int `json:"capacity"  validate:"required"`
}

type RequestTableUpdate struct {
	Number   int `json:"number"  validate:"required"`
	Capacity int `json:"capacity"  validate:"required"`
}
