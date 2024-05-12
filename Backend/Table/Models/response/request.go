package response

type RequestTableCrate struct {
	Number   int32 `json:"number" gorm:"type:int" validate:"required"`
	Capacity int32 `json:"capacity" gorm:"type:int" validate:"required"`
}

type RequestTableUpdate struct {
	Number   int32 `json:"number" gorm:"type:int" validate:"required"`
	Capacity int32 `json:"capacity" gorm:"type:int" validate:"required"`
}
