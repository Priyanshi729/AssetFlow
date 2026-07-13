package models

type RegisterUser struct {
	Name        string `json:"name" db:"name" validate:"required"`
	Email       string `json:"email" db:"email" validate:"required,email"`
	Password    string `json:"password" db:"password" validate:"required,min=6"`
	Role        string `json:"role" db:"role" validate:"required,oneof=admin employee project_manager"`
	UserType    string `json:"usertype" db:"user_type" validate:"required,oneof=full-time intern freelancer"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number" validate:"required,len=10"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	UserID       string `db:"user_id"`
	PasswordHash string `db:"password"`
}

type User struct {
	UserID      string `json:"id" db:"user_id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	Role        string `json:"role" db:"role"`
	UserType    string `json:"usertype" db:"user_type"`
}
