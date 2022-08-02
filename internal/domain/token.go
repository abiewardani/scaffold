package domain

import "github.com/dgrijalva/jwt-go"

type AccessTokenClaim struct {
	jwt.StandardClaims
	UserID   string `json:"user_id"`
	Username string `json:"user_name"`
}
