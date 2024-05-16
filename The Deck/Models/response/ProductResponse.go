package response

type ProductResponse struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Price       string `json:"price" validate:"required"`
	CategoryID  uint   `json:"category_id"`
}
