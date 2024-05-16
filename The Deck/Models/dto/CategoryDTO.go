package dto

type RequestCategoryCreate struct {
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
}
type RequestCategoryUpdate struct {
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
}
