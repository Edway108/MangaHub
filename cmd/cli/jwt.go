package main

import (
	"errors"

	"MangaHub/pkg/utils"
)

func getUserIDFromToken() (string, error) {
	token, err := loadToken()
	if err != nil {
		return "", err
	}

	claims, err := utils.ParseToken(token)
	if err != nil {
		return "", errors.New("invalid token, please login again")
	}

	return claims.UserID, nil
}
