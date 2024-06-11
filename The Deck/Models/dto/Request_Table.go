package dto

type AddRequestTable struct {
	TableId uint   `json:"table_id"`
	Notes   string `json:"notes"`
}

type UpdateRequestTable struct {
	Status int `json:"status"`
}
