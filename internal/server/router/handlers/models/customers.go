package models

type CustomerPayload struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email,max=50"`
	Phone     string `json:"phone" validate:"required,max=50"`
	Password  string `json:"password" validate:"required,min=6,max=20"`
}
