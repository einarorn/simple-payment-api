package app

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "rS4p0Xe1ZKvKZoBh17UzTF0p51UiK7TziwMKr9CtipLT6wquGsBC5AWzX8XgGzx"
const ValidMinutes = 5

func getToken(name string) (string, error) {
	signingKey := []byte(SecretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"exp": time.Now().Add(time.Minute * ValidMinutes).Unix(),
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(SecretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}