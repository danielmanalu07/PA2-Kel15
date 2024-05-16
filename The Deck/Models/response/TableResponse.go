package response

type TableResponse struct {
	Number   int `json:"number"  validate:"required"`
	Capacity int `json:"capacity"  validate:"required"`
}
