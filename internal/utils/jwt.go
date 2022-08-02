package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/abiewardani/scaffold/internal/domain"
)

type Token struct {
	Access  string
	Refresh string
}

func GenerateToken(UserID, ID, Username, Key string, AccessTokenLifeTime time.Duration) (*Token, error) {
	token := Token{}
	// GENERATE JWT TOKEN
	now := time.Now()
	end := now.Add(AccessTokenLifeTime)

	claim := domain.AccessTokenClaim{
		UserID:   UserID,
		Username: Username,
	}
	claim.IssuedAt = now.Unix()
	claim.ExpiresAt = end.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessTokenStr, err := newToken.SignedString([]byte(Key))
	if err != nil {
		return nil, err
	}
	token.Access = accessTokenStr

	rtClaims := domain.AccessTokenClaim{}
	rtClaims.IssuedAt = now.Unix()
	rtClaims.ExpiresAt = end.Add(AccessTokenLifeTime).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshTokenStr, err := refreshToken.SignedString([]byte(Key))
	if err != nil {
		return nil, err
	}
	token.Refresh = refreshTokenStr

	return &token, nil
}
