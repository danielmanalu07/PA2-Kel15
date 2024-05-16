package dto

type RequestProductCreate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       string `json:"price" validate:"required"`
	CategoryID  uint   `json:"category_id"`
}

type RequestProductUpdate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       string `json:"price" validate:"required"`
	CategoryID  uint   `json:"category_id"`
}
