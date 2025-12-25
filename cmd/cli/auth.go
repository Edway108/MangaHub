package main

import (
	"errors"
	"os"
	"strings"
)

func loadToken() (string, error) {
	data, err := os.ReadFile(".mangahub_token")
	if err != nil {
		return "", errors.New("not logged in, please run mangahub login")
	}

	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", errors.New("empty token, please login again")
	}

	return token, nil
}
