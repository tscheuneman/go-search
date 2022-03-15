package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/tscheuneman/go-search/container"
)

const defaultkey = "key"

func GenerateToken(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	tokenClaims := make(jwt.MapClaims)
	tokenClaims["user_id"] = user
	tokenClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token.Claims = tokenClaims

	tokenString, err := token.SignedString(container.JWT_KEY)

	return tokenString, err
}

func ValidateToken(jwtToken string, user string) bool {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		claims := token.Claims.(jwt.MapClaims)
		if claims["user_id"] != user {
			return nil, errors.New("Invalid User")
		}

		return []byte(container.JWT_KEY), nil
	})

	if err == nil && token.Valid {
		return true
	}
	return false
}
