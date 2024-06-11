package response

type TableResponse struct {
	Id       uint `json:"id"`
	Number   int  `json:"number"  validate:"required"`
	Capacity int  `json:"capacity"  validate:"required"`
	AdminID  uint `json:"admin_id"`
	Status   int  `json:"status"`
}
