package dto

type RequestOrderCreate struct {
	ProductIDs    []uint `json:"product_ids"`
	Total         string `json:"total"`
	Note          string `json:"note"`
	PaymentMethod string `json:"payment_method"`
	TableId       uint   `json:"table_id"`
	PickUpType    string `json:"pick_up_type"`
}

type RequestOrderUpdateStatus struct {
	Status int `json:"status"`
}
