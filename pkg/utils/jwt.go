package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("MQTM")

func GenerateToken(userID string, username string) (string, error) {
	//bassically this one is payload
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//add up sign
	return token.SignedString(SecretKey)
}
func ParseToken(tokenstring string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}) //call back to check secretkey
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
