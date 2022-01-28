package util

import (
	"errors"

	// "strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(email, name, userName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["name"] = name
	claims["handle"] = userName
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString([]byte(MySecretKeyForJWT))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateHeader(bearerHeader string) (interface{}, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerHeader, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		}
		return []byte(MySecretKeyForJWT), nil
	})
	if err != nil {
		return "", err
	}
	if token.Valid {
		return claims["handle"].(string), nil
	}
	return "", errors.New("invalid token")
}
