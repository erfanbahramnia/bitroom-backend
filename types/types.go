package types

import "github.com/dgrijalva/jwt-go"

type UserDataJwtClaims struct {
	Id    uint
	Phone string
	Role  string
}

type JwtCustomClaims struct {
	Id    uint
	Phone string
	Role  string
	jwt.StandardClaims
}

type JwtTokens struct {
	Token        string
	RefreshToken string
}

type ValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"mgs"`
}

type CustomError struct {
	Message string
	Code    int
}

type ExsitenceChecker func(uint) (bool, *CustomError)
