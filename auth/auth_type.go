package auth

import (
	"bitroom/types"
)

type RequiredDataForOtp struct {
	Phone string `json:"phone" validate:"required,min=11,max=11"`
}

type ValidateOtp struct {
	Phone string `json:"phone" validate:"required,min=11,max=11"`
	Otp   string `json:"otp" validate:"required,min=5,max=5"`
}

type RegisterResponse struct {
	UserId uint `json:"user_id"`
	types.JwtTokens
}

type LoginCredential struct {
	Phone    string `json:"phone" validate:"required,min=11,max=11"`
	Password string `json:"password" validate:"required,min=3,max=20"`
}

type AuthResponse struct {
	Phone      string          `json:"phone" validate:"required,min=11,max=11"`
	First_name string          `json:"first_name" validate:"required,min=3,max=30"`
	Last_name  string          `json:"last_name" validate:"required,min=3,max=40"`
	ID         uint            `json:"id"`
	Role       string          `json:"role"`
	Jwt        types.JwtTokens `json:"jwt"`
}
