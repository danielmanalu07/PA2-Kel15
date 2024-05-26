package response

type CustomerResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(100)" validate:"required"`
	Username    string `json:"username" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Email       string `json:"email" gorm:"type:varchar(100);unique" validate:"required,email"`
	Password    string `json:"password" gorm:"type:varchar(100)" validate:"required"`
	Phone       string `json:"phone" gorm:"type:varchar(100)" validate:"required"`
	Address     string `json:"address" gorm:"type:varchar(100)" validate:"required"`
	Gender      string `json:"gender" gorm:"type:varchar(50)" validate:"required"`
	DateOfBirth string `json:"date_of_birth" gorm:"type:varchar(100)" validate:"required"`
	Image       string `json:"image" gorm:"type:varchar(100)" validate:"required"`
}
