package types

import (
	"github.com/dgrijalva/jwt-go"
)

type LikeOrDislikeArticle struct {
	UserId    uint `json:"-"`
	ArticleId uint `json:"article_id" validate:"required"`
}

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

type ReactionChecker func(*LikeOrDislikeArticle) (bool, *CustomError)

type ReactionRemover func(*LikeOrDislikeArticle) *CustomError
