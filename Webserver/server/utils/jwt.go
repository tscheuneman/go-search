package utils

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/tscheuneman/go-search/container"
)

func GenerateToken(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	tokenClaims := make(jwt.MapClaims)
	tokenClaims["user_id"] = user
	tokenClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token.Claims = tokenClaims
	tokenString, err := token.SignedString([]byte(container.JWT_KEY))

	return tokenString, err
}

func ValidateToken(jwtToken string) (bool, string) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(container.JWT_KEY), nil
	})

	claims := token.Claims.(jwt.MapClaims)

	userId := claims["user_id"].(string)

	if userId == "" {
		return false, ""
	}

	if err == nil && token.Valid {
		return true, userId
	}

	return false, ""
}
