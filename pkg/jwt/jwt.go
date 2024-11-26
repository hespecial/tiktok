package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"tiktok/config"
	"time"
)

type Claims struct {
	UserId int64
	jwt.RegisteredClaims
}

func GenerateToken(userId int64) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.JwtIssuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * config.JwtTtl)),
		},
	}).SignedString(config.JwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
