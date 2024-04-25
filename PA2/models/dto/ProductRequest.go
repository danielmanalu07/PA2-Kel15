package dto

type RequestCreateProduct struct {
	Name        string  `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string  `json:"description" gorm:"type:text" validate:"required"`
	Price       float64 `json:"price" gorm:"type:float" validate:"required"`
	CategoryId  uint    `json:"category_id" gorm:"index"`
}

type RequestUpdateProduct struct {
	Name        string  `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string  `json:"description" gorm:"type:text" validate:"required"`
	Price       float64 `json:"price" gorm:"type:float" validate:"required"`
	CategoryId  uint    `json:"category_id" gorm:"index"`
}
