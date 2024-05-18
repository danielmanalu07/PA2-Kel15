package dto

type RequestCustomerRegister struct {
	FullName    string `json:"full_name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}

type RequestCustomerLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RequestCustomerUpdateProfile struct {
	FullName    string `json:"full_name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}

type RequestCustomerEditPassword struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestCustomerForgotPassword struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
