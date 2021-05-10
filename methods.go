package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "498e5outdrjijovlo5eijrnlioj"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
