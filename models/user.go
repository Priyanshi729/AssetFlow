package models

type RegisterUser struct {
	Name        string `json:"name" db:"name" validate:"required"`
	Email       string `json:"email" db:"email" validate:"required,email"`
	Password    string `json:"password" db:"password" validate:"required,min=6"`
	Role        string `json:"role" db:"role" validate:"required,oneof=admin employee project_manager"`
	Type        string `json:"type" db:"type" validate:"required,oneof=full_time intern freelancer"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number" validate:"required,len=10"`
}

type LoginUser struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=6"`
}

type User struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	Role        string `json:"role" db:"role"`
	Employment  string `json:"employment" db:"employment"`
}
