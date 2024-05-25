package response

type TableResponse struct {
	Id       uint   `json:"id"`
	Number   int    `json:"number"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`
}
