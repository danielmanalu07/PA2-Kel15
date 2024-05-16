package response

type CategoryResponse struct {
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
}
