package utils

import (
	"bitroom/config"
	"bitroom/types"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(claim types.UserDataJwtClaims) (string, error) {
	tokenClaims := types.JwtCustomClaims{
		Id:    claim.Id,
		Phone: claim.Phone,
		Role:  claim.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(config.ServerConfig.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateJwtRefreshToken(claim types.UserDataJwtClaims) (string, error) {
	refreshClaims := types.JwtCustomClaims{
		Id:    claim.Id,
		Phone: claim.Phone,
		Role:  claim.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(config.ServerConfig.Jwt_REFRESH_SECRET))
	if err != nil {
		return "", err
	}
	return refreshString, nil
}

func GenerateJwt(claim types.UserDataJwtClaims) (*types.JwtTokens, error) {
	// generate access token
	token, err := GenerateJwtToken(claim)
	if err != nil {
		return nil, err
	}

	// generate referesh token
	refreshToken, err := GenerateJwtRefreshToken(claim)
	if err != nil {
		return nil, err
	}

	// get all jwt tokens
	jwt := &types.JwtTokens{
		Token:        token,
		RefreshToken: refreshToken,
	}
	return jwt, nil
}
