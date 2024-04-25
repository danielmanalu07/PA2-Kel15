package dto

type RequestAdminCreate struct {
	Username string `json:"username" gorm:"type:varchar(255)" validate:"required,min=5"`
	Password string `json:"password" gorm:"type:varchar(255)" validate:"required,min=5"`
}

type RequestAdminLogin struct {
	Username string `json:"username" gorm:"type:varchar(255)" validate:"required,min=5"`
	Password string `json:"password" gorm:"type:varchar(255)" validate:"required,min=5"`
}
